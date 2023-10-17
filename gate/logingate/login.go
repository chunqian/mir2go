// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"fmt"
	"net"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"

	. "github.com/chunqian/mir2go/common"
	"github.com/chunqian/mir2go/gate/logingate/widget"
	log "github.com/chunqian/tinylog"
)

type TFrmMain struct {
	*vcl.TForm

	MainMenu  *widget.TMainMenu
	MemoLog   *vcl.TMemo
	Panel     *widget.TPanel
	StatusBar *vcl.TStatusBar

	SendTimer   *vcl.TTimer
	StartTimer  *vcl.TTimer
	DecodeTimer *vcl.TTimer
	Timer       *vcl.TTimer

	ClientSocket *TClientSocket
	ServerSocket *TServerSocket
}

// ******************** Var ********************
var (
	frmMain *TFrmMain
)

// ******************** TFrmMain ********************
func (sf *TFrmMain) SetComponents() {

	sf.MainMenu = widget.NewMainMenu(sf)

	sf.Panel = widget.NewPanel(sf)
	sf.Panel.SetAlign(types.AlTop)
	sf.Panel.SetBevelOuter(types.BvNone)
	sf.Panel.SetTabOrder(1)
	sf.Panel.SetBounds(0, 0, 308, 119)

	sf.MemoLog = vcl.NewMemo(sf)
	sf.MemoLog.SetName("MemoLog")
	sf.MemoLog.SetAlign(types.AlClient)
	sf.MemoLog.SetText("")
	sf.MemoLog.SetColor(colors.ClMenuText)
	sf.MemoLog.Font().SetColor(colors.ClLimegreen)
	sf.MemoLog.SetBounds(0, 119, 308, 18)
	sf.MemoLog.SetWordWrap(false)
	sf.MemoLog.SetScrollBars(types.SsHorizontal)
	sf.MemoLog.SetReadOnly(true)

	sf.StatusBar = vcl.NewStatusBar(sf)
	sf.StatusBar.SetSimplePanel(false)
	sf.StatusBar.SetBounds(0, 137, 308, 17)
	spanel := sf.StatusBar.Panels().Add()
	spanel.SetAlignment(types.TaCenter)
	spanel.SetText("7100")
	spanel.SetWidth(50)
	spanel = sf.StatusBar.Panels().Add()
	spanel.SetAlignment(types.TaCenter)
	spanel.SetText("未连接")
	spanel.SetWidth(60)
	spanel = sf.StatusBar.Panels().Add()
	spanel.SetAlignment(types.TaCenter)
	spanel.SetText("0/0")
	spanel.SetWidth(70)
	spanel = sf.StatusBar.Panels().Add()
	spanel.SetWidth(50)

	sf.StartTimer = vcl.NewTimer(sf)
	sf.StartTimer.SetInterval(200)
	sf.StartTimer.SetEnabled(true)
	sf.StartTimer.SetOnTimer(sf.StartTimerTimer)

	sf.DecodeTimer = vcl.NewTimer(sf)
	sf.DecodeTimer.SetInterval(3)
	sf.DecodeTimer.SetEnabled(false)
	sf.DecodeTimer.SetOnTimer(sf.DecodeTimerTimer)

	sf.SendTimer = vcl.NewTimer(sf)
	sf.SendTimer.SetInterval(3000)
	sf.SendTimer.SetEnabled(false)
	sf.SendTimer.SetOnTimer(sf.SendTimerTimer)

	sf.Timer = vcl.NewTimer(sf)
	sf.Timer.SetInterval(1000)
	sf.Timer.SetEnabled(false)
	sf.Timer.SetOnTimer(sf.TimerTimer)

	sf.Panel.SetParent(sf)
	sf.MemoLog.SetParent(sf)
	sf.StatusBar.SetParent(sf)
}

func (sf *TFrmMain) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.SetCaption("登陆网关")
	sf.EnabledMaximize(false)
	sf.SetBorderStyle(types.BsSingle)
	sf.SetBounds(636, 215, 308, 154)
	sf.SetComponents()

	// 注册观察者
	GetSubject("TFrmMain").Register(frmMain)

	DecodeMsgTime = 0
	sf.initUserSessionArray()
}

func (sf *TFrmMain) OnFormDestroy(sender vcl.IObject) {
	for i := range SessionArray {
		SessionArray[i] = nil
	}
}

func (sf *TFrmMain) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	if Closed {
		return
	}
	*canClose = vcl.MessageDlg("是否确认退出服务器?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.MrYes
}

func (sf *TFrmMain) closeSocket(socketHandle uintptr) {
	for i := range SessionArray {
		userSession := SessionArray[i]
		if userSession.Socket != nil && userSession.SocketHandle == socketHandle {
			userSession.Socket.Close()
			break
		}
	}
}

func (sf *TFrmMain) initUserSessionArray() {
	for i := 0; i < GATEMAXSESSION; i++ {
		userSession := TUserSession{}
		userSession.Socket = nil
		userSession.RemoteIPAddr = ""
		userSession.SendMsgLen = 0
		userSession.SendLock = false
		userSession.SendAvailable = true
		userSession.SendCheck = false
		userSession.CheckSendLength = 0
		userSession.ReceiveLength = 0
		userSession.UserTimeOutTick = GetTickCount()
		userSession.SocketHandle = uintptr(0)
		userSession.MsgList = make([]string, 0)
		SessionArray[i] = &userSession
	}
}

func (sf *TFrmMain) isBlockIP(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	for i := range TempBlockIPList {
		sockAddr := TempBlockIPList[i]
		if sockAddr.IP == int(IpToInteger(ip)) {
			return true
		}
	}
	for i := range BlockIPList {
		sockAddr := BlockIPList[i]
		if sockAddr.IP == int(IpToInteger(ip)) {
			return true
		}
	}
	return false
}

func (sf *TFrmMain) isConnLimited(ipAddr string) bool {
	denyConnect := false
	ip := net.ParseIP(ipAddr)
	for i := range CurrIPAddrList {
		sockAddr := CurrIPAddrList[i]
		if sockAddr.IP == int(IpToInteger(ip)) {
			sockAddr.Count++
			if (GetTickCount() - sockAddr.IPCountTick1) < 1000 {
				sockAddr.IPCount1++
				if sockAddr.IPCount1 >= IPCountLimit1 {
					denyConnect = true
				}
			} else {
				sockAddr.IPCountTick1 = GetTickCount()
				sockAddr.IPCount1 = 0
			}

			if (GetTickCount() - sockAddr.IPCountTick2) < 3000 {
				sockAddr.IPCount2++
				if sockAddr.IPCount2 >= IPCountLimit2 {
					denyConnect = true
				}
			} else {
				sockAddr.IPCountTick2 = GetTickCount()
				sockAddr.IPCount2 = 0
			}

			if sockAddr.Count > int(MaxConnOfIPAddr) {
				denyConnect = true
			}
		}
	}

	sockAddr := TSockAddr{
		IP:    int(IpToInteger(ip)),
		Count: 1,
	}
	CurrIPAddrList = append(CurrIPAddrList, sockAddr)
	return denyConnect
}

func (sf *TFrmMain) loadConfig() {
	conf := vcl.NewIniFile(ConfigFile)
	if conf != nil {
		TitleName = conf.ReadString(GateClass, "Title", TitleName)
		ServerPort = conf.ReadInteger(GateClass, "ServerPort", ServerPort)
		ServerAddr = conf.ReadString(GateClass, "ServerAddr", ServerAddr)
		GatePort = conf.ReadInteger(GateClass, "GatePort", GatePort)
		GateAddr = conf.ReadString(GateClass, "GateAddr", GateAddr)
		MainLog.SetLogLevel(conf.ReadInteger(GateClass, "ShowLogLevel", 3))

		BlockMethod = TBlockIPMethod(conf.ReadInteger(GateClass, "BlockMethod", int32(BlockMethod)))

		if conf.ReadInteger(GateClass, "KeepConnectTimeOut", -1) <= 0 {
			conf.WriteInteger(GateClass, "KeepConnectTimeOut", KeepConnectTimeOut)
		}

		MaxConnOfIPAddr = conf.ReadInteger(GateClass, "MaxConnOfIPAddr", MaxConnOfIPAddr)
		KeepConnectTimeOut = conf.ReadInteger(GateClass, "KeepConnectTimeOut", KeepConnectTimeOut)
		DynamicIPDisMode = conf.ReadBool(GateClass, "DynamicIPDisMode", DynamicIPDisMode)

		conf.Free()
	}

	LoadBlockIPFile()
}

func (sf *TFrmMain) resUserSessionArray() {
	for i := 0; i < GATEMAXSESSION; i++ {
		userSession := SessionArray[i]
		userSession.Socket = nil
		userSession.RemoteIPAddr = ""
		userSession.SocketHandle = uintptr(0)
		userSession.MsgList = userSession.MsgList[:0]
	}
}

func (sf *TFrmMain) sendUserMsg(userSession *TUserSession, sendMsg string) int {
	result := -1
	if userSession.Socket != nil {
		if !userSession.SendLock {
			if userSession.SendAvailable && GetTickCount() > userSession.SendLockTimeOut {
				userSession.SendAvailable = true
				userSession.CheckSendLength = 0
				SendHoldTimeOut = true
				SendHoldTick = GetTickCount()
			}
			if userSession.SendAvailable {
				if userSession.CheckSendLength >= 250 {
					if !userSession.SendCheck {
						userSession.SendCheck = true
						sendMsg = "*" + sendMsg
					}
					if userSession.CheckSendLength >= 512 {
						userSession.SendAvailable = false
						userSession.SendLockTimeOut = GetTickCount() + 3*1000
					}
				}
				userSession.Socket.Write([]byte(sendMsg))
				userSession.SendMsgLen += len(sendMsg)
				userSession.CheckSendLength += len(sendMsg)
				result = 1
			}
		} else {
			result = 0
		}
	} else {
		result = 0
	}
	return result
}

func (sf *TFrmMain) showLogMsg(flag bool) {
	var height int32
	if flag {
		height = sf.Panel.Height()
		sf.Panel.SetHeight(0)
		sf.MemoLog.SetHeight(height)
		sf.MemoLog.SetTop(sf.Panel.Top())
	} else {
		height = sf.MemoLog.Height()
		sf.MemoLog.SetHeight(0)
		sf.Panel.SetHeight(height)
	}
}

func (sf *TFrmMain) showMainLogMsg() {

	if GetTickCount()-ShowMainLogTick < 200 {
		return
	}
	ShowMainLogTick = GetTickCount()

	// 获取主日志列表, 在 GUI 中显示日志
	memoLog := vcl.AsMemo(sf.FindComponent("MemoLog"))

	// 更新UI
	for _, logMsg := range MainLog.MsgList() {
		memoLog.Lines().Add(logMsg)
	}

	// 清空主日志列表
	MainLog.ClearMsgList()
}

func (sf *TFrmMain) startService() {

	defer func() {
		if r := recover(); r != nil {
			// 通知
			GetSubject("widget.TMainMenu.TMenuControl").Notify("SetMenuControlStart", true)
			GetSubject("widget.TMainMenu.TMenuControl").Notify("SetMenuControlStop", false)
			MainLog.AddLogMsg(fmt.Sprintf("%v", r), 0)
		}
	}()

	// 在这里添加启动服务的逻辑
	// 初始化变量和状态
	MainLog.AddLogMsg("正在启动服务...", 3)

	ServiceStart = true
	GateReady = true
	ServerReady = false
	SessionCount = 0
	// 通知
	GetSubject("widget.TMainMenu.TMenuControl").Notify("SetMenuControlStart", false)
	GetSubject("widget.TMainMenu.TMenuControl").Notify("SetMenuControlStop", true)

	ReConnectServerTick = GetTickCount() - 25*1000
	KeepAliveTimeOut = false
	SendMsgCount = 0
	TotalMsgListCount = 0
	SendKeepAliveTick = GetTickCount()
	SendHoldTimeOut = false
	SendHoldTick = GetTickCount()

	CurrIPAddrList = make([]TSockAddr, 0)
	BlockIPList = make([]TSockAddr, 0)
	TempBlockIPList = make([]TSockAddr, 0)
	ClientSockeMsgList = make([]string, 0)

	sf.resUserSessionArray()
	sf.loadConfig()

	// 更新UI
	if TitleName != "" {
		sf.SetCaption(GateName + " - " + TitleName)
	} else {
		sf.SetCaption(GateName)
	}

	sf.ClientSocket = &TClientSocket{}
	sf.ClientSocket.Dial(sf, ServerAddr, ServerPort)

	sf.ServerSocket = &TServerSocket{}
	sf.ServerSocket.Listen(sf, GateAddr, GatePort)

	sf.DecodeTimer.SetEnabled(true)
	sf.SendTimer.SetEnabled(true)
	sf.Timer.SetEnabled(true)

	MainLog.AddLogMsg("启动服务完成...", 3)
}

func (sf *TFrmMain) stopService() {

	MainLog.AddLogMsg("正在停止服务...", 3)
	ServiceStart = false
	GateReady = false
	// 通知
	GetSubject("widget.TMainMenu.TMenuControl").Notify("SetMenuControlStart", true)
	GetSubject("widget.TMainMenu.TMenuControl").Notify("SetMenuControlStop", false)

	sf.SendTimer.SetEnabled(false)

	for i := 0; i < GATEMAXSESSION; i++ {
		if SessionArray[i].Socket != nil {
			SessionArray[i].Socket.Close()
		}
	}

	SaveBlockIPList()

	sf.ServerSocket.Close()
	sf.ClientSocket.Close()

	CurrIPAddrList = CurrIPAddrList[:0]
	BlockIPList = BlockIPList[:0]
	TempBlockIPList = TempBlockIPList[:0]
	ClientSockeMsgList = ClientSockeMsgList[:0]

	MainLog.AddLogMsg("停止服务完成...", 3)
}

func (sf *TFrmMain) CloseConnect(ipAddr string) {
	if sf.ServerSocket.Active() {
		for {
			check := false
			for i := 0; i < sf.ServerSocket.ActiveConnections(); i++ {
				remoteAddr := GetAddrHost(sf.ServerSocket.Connections()[i].RemoteAddr())
				if ipAddr == remoteAddr {
					sf.ServerSocket.Connections()[i].Close()
					check = true
					break
				}
			}
			if !check {
				break
			}
		}
	}
}

func (sf *TFrmMain) ClientSocketConnect(socket *TClientSocket) {
	GateReady = true
	SessionCount = 0
	KeepAliveTick = GetTickCount()
	sf.resUserSessionArray()
	ServerReady = true
}

func (sf *TFrmMain) ClientSocketDisconnect(socket *TClientSocket, err error) {
	log.Info("ClientSocketDisconnect: {}", err.Error())
	for i := 0; i < GATEMAXSESSION; i++ {
		userSession := SessionArray[i]
		if userSession.Socket != nil {
			userSession.Socket.Close()
		}
		userSession.Socket = nil
		userSession.RemoteIPAddr = ""
		userSession.SocketHandle = uintptr(0)
	}

	sf.resUserSessionArray()
	ClientSockeMsgList = ClientSockeMsgList[:0]
	GateReady = false
	SessionCount = 0
}

func (sf *TFrmMain) ClientSocketError(socket *TClientSocket, err error) {
	log.Info("ClientSocketError: {}", err.Error())
	socket.Close()
	ServerReady = false
}

func (sf *TFrmMain) ClientSocketRead(socket *TClientSocket, message string) {
	log.Info("ClientSocketRead: {}", message)
	ClientSockeMsgList = append(ClientSockeMsgList, message)
}

func (sf *TFrmMain) DecodeTimerTimer(sender vcl.IObject) {
	var processMsg, socketMsg string

	sf.showMainLogMsg()

	if DecodeLock || !GateReady {
		return
	}

	decodeTick := GetTickCount()
	DecodeLock = true
	defer func() {
		DecodeLock = false
	}()
	processMsg = ""

	for {
		if len(ClientSockeMsgList) <= 0 {
			break
		}
		processMsg = ProcMsg + ClientSockeMsgList[0]
		ProcMsg = ""
		ClientSockeMsgList = ClientSockeMsgList[1:]
		for {
			if TagCount(processMsg, '$') < 1 {
				break
			}
			processMsg, socketMsg = ArrestStringEx(processMsg, "%", "$")
			log.Info("processMsg: {}, socketMsg: {}", processMsg, socketMsg)
			if socketMsg == "" {
				break
			}
			if socketMsg[0] == '+' {
				if socketMsg[1] == '-' {
					sf.closeSocket(
						uintptr(StrToInt(socketMsg[2:len(socketMsg)-2], 0)),
					)
					continue
				} else {
					KeepAliveTick = GetTickCount()
					KeepAliveTimeOut = false
					continue
				}
			}
			socketMsg, socketHandleStr := GetValidStr3(socketMsg, []rune{'/'})
			socketHandle := uintptr(StrToInt(socketHandleStr, 0))
			if socketHandle <= 0 {
				continue
			}
			for i := 0; i < GATEMAXSESSION; i++ {
				if SessionArray[i].Socket.SocketHandle() == socketHandle {
					SessionArray[i].MsgList = append(SessionArray[i].MsgList, socketMsg)
					break
				}
			}
		}
	}

	if processMsg != "" {
		ProcMsg = processMsg
	}

	SendMsgCount = 0
	TotalMsgListCount = 0

	for i := 0; i < GATEMAXSESSION; i++ {
		if SessionArray[i].SocketHandle <= 0 {
			continue
		}
		// 踢除超时无数据传输连接
		if GetTickCount()-SessionArray[i].ConnectCheckTick > uint32(KeepConnectTimeOut) {
			remoteIPAddr := SessionArray[i].RemoteIPAddr
			switch BlockMethod {
			case Disconnect:
				SessionArray[i].Socket.Close()
			case Block:
				ip := net.ParseIP(remoteIPAddr)
				sockAddr := TSockAddr{}
				sockAddr.IP = int(IpToInteger(ip))
				TempBlockIPList = append(TempBlockIPList, sockAddr)
				sf.CloseConnect(remoteIPAddr)
			case BlockList:
				ip := net.ParseIP(remoteIPAddr)
				sockAddr := TSockAddr{}
				sockAddr.IP = int(IpToInteger(ip))
				BlockIPList = append(BlockIPList, sockAddr)
				sf.CloseConnect(remoteIPAddr)
			}
			MainLog.AddLogMsg("端口空连接攻击: "+remoteIPAddr, 1)
			continue
		}

		for {
			if len(SessionArray[i].MsgList) <= 0 {
				break
			}
			userSession := SessionArray[i]

			sendRetCode := sf.sendUserMsg(userSession, userSession.MsgList[0])
			if sendRetCode >= 0 {
				if sendRetCode == 1 {
					userSession.ConnectCheckTick = GetTickCount()
					userSession.MsgList = userSession.MsgList[1:]
					continue
				}
				if len(userSession.MsgList) > 100 {
					msgCount := 0
					for msgCount != 51 {
						userSession.MsgList = userSession.MsgList[1:]
						msgCount++
					}
				}
				TotalMsgListCount += len(userSession.MsgList)
				MainLog.AddLogMsg(userSession.RemoteIPAddr+" : "+IntToStr(len(userSession.MsgList)), 5)
				SendMsgCount++
			} else {
				userSession.SocketHandle = uintptr(0)
				userSession.Socket = nil
				userSession.MsgList = userSession.MsgList[:0]
			}
		}
	}

	if GetTickCount()-SendKeepAliveTick > 2*1000 {
		SendKeepAliveTick = GetTickCount()
		if GateReady {
			sf.ClientSocket.Write([]byte("%--$"))
		}
	}

	if GetTickCount()-KeepAliveTick > 10*1000 {
		KeepAliveTimeOut = true
		sf.ClientSocket.Close()
	}

	decodeTime := GetTickCount() - decodeTick
	if DecodeMsgTime < decodeTime {
		DecodeMsgTime = decodeTime
	}
	if DecodeMsgTime > 50 {
		DecodeMsgTime -= 50
	}
}

func (sf *TFrmMain) MemoLogChange(sender vcl.IObject) {
	if sf.MemoLog.Lines().Count() > 200 {
		sf.MemoLog.Clear()
	}
}

func (sf *TFrmMain) SendTimerTimer(sender vcl.IObject) {
	if sf.ServerSocket.Active() {
		ActiveConnections = sf.ServerSocket.ActiveConnections()
	}

	// 更新UI
	if SendHoldTimeOut {
		sf.Panel.Hold.SetCaption(IntToStr(ActiveConnections) + "#")
		if GetTickCount()-SendHoldTick > 3*1000 {
			SendHoldTimeOut = false
		}
	} else {
		sf.Panel.Hold.SetCaption(IntToStr(ActiveConnections))
	}

	if GateReady && !KeepAliveTimeOut {
		for i := 0; i < GATEMAXSESSION; i++ {
			userSession := SessionArray[i]
			if userSession.Socket != nil {
				if GetTickCount()-userSession.UserTimeOutTick > 60*60*1000 {
					userSession.Socket.Close()
					userSession.Socket = nil
					userSession.SocketHandle = uintptr(0)
					userSession.MsgList = userSession.MsgList[:0]
					userSession.RemoteIPAddr = ""
				}
			}
		}
	}
	if !GateReady && ServiceStart {
		if GetTickCount()-ReConnectServerTick > 1000 /* 30*1000 */ {
			ReConnectServerTick = GetTickCount()
			sf.ClientSocket.Dial(sf, ServerAddr, ServerPort)
		}
	}
}

func (sf *TFrmMain) ServerSocketClientConnect(socket *TClientSocket) {
	var remoteIPAddr, localIPAddr string
	var sockAddr TSockAddr

	socket.Index = -1
	remoteIPAddr = GetAddrHost(socket.RemoteAddr())

	if DynamicIPDisMode {
		localIPAddr = GetAddrHost(sf.ClientSocket.RemoteAddr())
	} else {
		localIPAddr = GetAddrHost(socket.LocalAddr())
	}

	if sf.isBlockIP(remoteIPAddr) {
		MainLog.AddLogMsg("过滤连接: "+remoteIPAddr, 1)
		socket.Close()
		return
	}

	if sf.isConnLimited(remoteIPAddr) {
		switch BlockMethod {
		case Disconnect:
			socket.Close()
		case Block:
			ip := net.ParseIP(remoteIPAddr)
			sockAddr = TSockAddr{}
			sockAddr.IP = int(IpToInteger(ip))
			TempBlockIPList = append(TempBlockIPList, sockAddr)
			sf.CloseConnect(remoteIPAddr)
		case BlockList:
			ip := net.ParseIP(remoteIPAddr)
			sockAddr = TSockAddr{}
			sockAddr.IP = int(IpToInteger(ip))
			BlockIPList = append(BlockIPList, sockAddr)
			sf.CloseConnect(remoteIPAddr)
		}
		MainLog.AddLogMsg("端口攻击: "+remoteIPAddr, 1)
	}

	if GateReady {
		for i := 0; i < GATEMAXSESSION; i++ {
			userSession := SessionArray[i]
			if userSession.Socket == nil {
				userSession.Socket = socket
				userSession.RemoteIPAddr = remoteIPAddr
				userSession.SendMsgLen = 0
				userSession.SendLock = false
				userSession.ConnectCheckTick = GetTickCount()
				userSession.SendAvailable = true
				userSession.SendCheck = false
				userSession.CheckSendLength = 0
				userSession.ReceiveLength = 0
				userSession.UserTimeOutTick = GetTickCount()
				userSession.SocketHandle = socket.SocketHandle()
				userSession.MsgList = make([]string, 0)

				socket.Index = i
				SessionCount++
				break
			}
		}

		if socket.Index >= 0 {
			// 发送字符串
			message := fmt.Sprintf("%%O%d/%s/%s$", socket.SocketHandle(), remoteIPAddr, localIPAddr)
			sf.ClientSocket.Write([]byte(message))
			MainLog.AddLogMsg("Connect: "+remoteIPAddr, 5)
		} else {
			socket.Close()
			MainLog.AddLogMsg("Kick Off: "+remoteIPAddr, 1)
		}
	} else {
		socket.Close()
		MainLog.AddLogMsg("Kick Off: "+remoteIPAddr, 1)
	}
}

func (sf *TFrmMain) ServerSocketClientDisconnect(conn *TClientSocket, err error) {
	log.Info("ServerSocketClientDisconnect: {}", err.Error())
	remoteIPAddr := GetAddrHost(conn.RemoteAddr())
	ip := net.ParseIP(remoteIPAddr)
	sockIndex := conn.Index

	for i := 0; i < len(CurrIPAddrList); i++ {
		sockAddr := CurrIPAddrList[i]
		if sockAddr.IP == int(IpToInteger(ip)) {
			sockAddr.Count--
			if sockAddr.Count <= 0 {
				CurrIPAddrList = append(CurrIPAddrList[:i], CurrIPAddrList[i+1:]...)
			}
			break
		}
	}

	if sockIndex >= 0 && sockIndex < GATEMAXSESSION {
		userSession := SessionArray[sockIndex]
		userSession.Socket = nil
		userSession.RemoteIPAddr = ""
		userSession.SocketHandle = uintptr(0)
		userSession.MsgList = userSession.MsgList[:0]
		SessionCount--
		if GateReady {
			message := fmt.Sprintf("%%X%d$", conn.SocketHandle())
			sf.ClientSocket.Write([]byte(message))
			MainLog.AddLogMsg("DisConnect: "+remoteIPAddr, 5)
		}
	}
}

func (sf *TFrmMain) ServerSocketClientError(conn *TClientSocket, err error) {
	log.Info("ServerSocketClientError: {}", err.Error())
	conn.Close()
}

func (sf *TFrmMain) ServerSocketClientRead(conn *TClientSocket, message string) {
	log.Info("ServerSocketClientRead: {}", message)
	sockIndex := conn.Index
	if sockIndex >= 0 && sockIndex < GATEMAXSESSION {
		userSession := SessionArray[sockIndex]
		if ServerReady {
			userSession.SendAvailable = true
			userSession.SendCheck = false
			userSession.CheckSendLength = 0
			if GateReady && !KeepAliveTimeOut {
				userSession.ConnectCheckTick = GetTickCount()
				if (GetTickCount() - userSession.UserTimeOutTick) < 1000 {
					userSession.ReceiveLength += len(message)
				} else {
					userSession.ReceiveLength = len(message)
				}
				message := fmt.Sprintf("%%A%d/%s$", conn.SocketHandle(), message)
				sf.ClientSocket.Write([]byte(message))
			}
		}
	}
}

func (sf *TFrmMain) StartTimerTimer(sender vcl.IObject) {

	startTimer := vcl.AsTimer(sender)
	if Started {
		startTimer.SetEnabled(false) // 禁用定时器
		sf.stopService()
		Closed = true
		vcl.Application.Terminate() // 关闭应用程序
	} else {
		// 通知
		GetSubject("widget.TMainMenu.TMenuView").Notify("MenuViewLogMsgClick", nil)
		Started = true
		startTimer.SetEnabled(false) // 禁用定时器
		sf.startService()
	}
}

func (sf *TFrmMain) TimerTimer(sender vcl.IObject) {
	var port string

	// 更新UI
	if sf.ServerSocket.Active() {
		port = GetAddrPort(sf.ServerSocket.Addr())
		sf.StatusBar.Panels().Items(0).SetText(port)
		if SendHoldTimeOut {
			sf.StatusBar.Panels().Items(2).SetText(
				IntToStr(SessionCount) + "/#" + IntToStr(sf.ServerSocket.ActiveConnections()),
			)
		} else {
			sf.StatusBar.Panels().Items(2).SetText(
				IntToStr(SessionCount) + "/" + IntToStr(sf.ServerSocket.ActiveConnections()),
			)
		}
	} else {
		sf.StatusBar.Panels().Items(0).SetText("????")
		sf.StatusBar.Panels().Items(2).SetText("????")
	}
	sf.Panel.Label2.SetCaption(IntToStr(int(DecodeMsgTime)))
	if !GateReady {
		sf.StatusBar.Panels().Items(1).SetText("未连接")
	} else {
		if KeepAliveTimeOut {
			sf.StatusBar.Panels().Items(1).SetText("超时")
		} else {
			sf.StatusBar.Panels().Items(1).SetText("已连接")
			sf.Panel.Lack.SetCaption(IntToStr(TotalMsgListCount) + "/" + IntToStr(SendMsgCount))
		}
	}
}

func (sf *TFrmMain) ObserverNotifyReceived(tag string, data interface{}) {
	switch tag {
	case "menuControlStartClick":
		sf.startService()
	case "menuControlStopClick":
		sf.stopService()
	case "menuControlReconnectClick":
		ReConnectServerTick = 0
	case "menuControlClearLogClick":
		sf.MemoLog.Clear()
	case "menuControlExitClick":
		sf.Close()
	case "menuViewLogMsgClick":
		sf.showLogMsg(data.(bool))
	}
}
