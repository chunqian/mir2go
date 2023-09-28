// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"

	"github.com/chunqian/mir2go/common"
	log "github.com/chunqian/tinylog"
)

// ******************** Const ********************
const (
	GATEMAXSESSION = 10000
)

type TFrmMain struct {
	*vcl.TForm

	MainMenu  *TMainMenu
	MemoLog   *vcl.TMemo
	Panel     *TPanel
	StatusBar *vcl.TStatusBar

	SendTimer    *vcl.TTimer
	StartTimer   *vcl.TTimer
	DecodeTimer  *vcl.TTimer
	Timer        *vcl.TTimer
	ClientSocket *TClientSocket
	ServerSocket *TServerSocket

	serverReady         bool
	showLocked          bool
	decodeMsgTime       uint32
	reConnectServerTick uint32
	sendKeepAliveTick   uint32
	showMainLogTick     uint32
	sessionCount        int
	tempLogList         []string
}

type TMainMenu struct {
	*vcl.TMainMenu

	MenuControl *TMenuControl
	MenuView    *TMenuView
	MenuOption  *TMenuOption
	N3          *TMenuItem3
}

type TPanel struct {
	*vcl.TPanel

	Label2 *vcl.TLabel
	Hold   *vcl.TLabel
	Lack   *vcl.TLabel
}

type TMenuControl struct {
	*vcl.TMenuItem

	MenuControlStart     *vcl.TMenuItem
	MenuControlStop      *vcl.TMenuItem
	MenuControlReconnect *vcl.TMenuItem
	N1                   *vcl.TMenuItem
	MenuControlClearLog  *vcl.TMenuItem
	N2                   *vcl.TMenuItem
	MenuControlExit      *vcl.TMenuItem
}

type TMenuView struct {
	*vcl.TMenuItem

	MenuViewLogMsg *vcl.TMenuItem
}

type TMenuOption struct {
	*vcl.TMenuItem

	MenuOptionGeneral  *vcl.TMenuItem
	MenuOptionIpFilter *vcl.TMenuItem
}

type TMenuItem3 struct {
	*vcl.TMenuItem

	N4 *vcl.TMenuItem
}

type TUserSession struct {
	Socket          *TClientSocket
	RemoteIPaddr    string
	SendMsgLen      int
	SendLock        bool
	CheckSendLength int
	SendAvailable   bool
	SendCheck       bool
	SendLockTimeOut uint32
	ReceiveLength   int
	UserTimeOutTick uint32
	SocketHandle    uintptr
	IP              string
	MsgList         []string
	MsgListMutex    sync.Mutex
	ConnctCheckTick uint32
}

type TSessionArray [GATEMAXSESSION]*TUserSession

// ******************** Var ********************
var (
	ClientSockeMsgList []string
	FrmMain            *TFrmMain
	SessionArray       TSessionArray
	ProcMsg            string
)

// ******************** Layout ********************
func (sf *TFrmMain) Layout() {

	sf.MainMenu = &TMainMenu{
		TMainMenu: vcl.NewMainMenu(sf),
	}
	sf.Panel = &TPanel{
		TPanel: vcl.NewPanel(sf),
	}

	sf.MemoLog = vcl.NewMemo(sf)
	sf.MemoLog.SetName("MemoLog")
	sf.MemoLog.SetAlign(types.AlClient)
	sf.MemoLog.SetText("")
	sf.MemoLog.SetParent(sf)
	sf.MemoLog.SetColor(colors.ClMenuText)
	sf.MemoLog.Font().SetColor(colors.ClLimegreen)
	sf.MemoLog.SetLeft(0)
	sf.MemoLog.SetTop(119)
	sf.MemoLog.SetWidth(308)
	sf.MemoLog.SetHeight(18)
	sf.MemoLog.SetWordWrap(false)
	sf.MemoLog.SetScrollBars(types.SsHorizontal)

	sf.StatusBar = vcl.NewStatusBar(sf)
	sf.StatusBar.SetParent(sf)
	sf.StatusBar.SetSimplePanel(false)
	sf.StatusBar.SetLeft(0)
	sf.StatusBar.SetTop(137)
	sf.StatusBar.SetWidth(308)
	sf.StatusBar.SetHeight(17)
	panel := sf.StatusBar.Panels().Add()
	panel.SetAlignment(types.TaCenter)
	panel.SetText("7100")
	panel.SetWidth(50)
	panel = sf.StatusBar.Panels().Add()
	panel.SetAlignment(types.TaCenter)
	panel.SetText("未连接")
	panel.SetWidth(60)
	panel = sf.StatusBar.Panels().Add()
	panel.SetAlignment(types.TaCenter)
	panel.SetText("0/0")
	panel.SetWidth(70)
	panel = sf.StatusBar.Panels().Add()
	panel.SetWidth(50)

	sf.Panel.Layout(sf)
	sf.MainMenu.Layout(sf)

	sf.StartTimer = vcl.NewTimer(sf)
	sf.StartTimer.SetInterval(200)
	sf.StartTimer.SetEnabled(true)
	sf.StartTimer.SetOnTimer(sf.StartTimerTimer)

	sf.DecodeTimer = vcl.NewTimer(sf)
	sf.DecodeTimer.SetInterval(1)
	sf.DecodeTimer.SetEnabled(false)
	sf.DecodeTimer.SetOnTimer(sf.DecodeTimerTimer)

	sf.SendTimer = vcl.NewTimer(sf)
	sf.SendTimer.SetInterval(3000)
	sf.SendTimer.SetEnabled(false)
	sf.SendTimer.SetOnTimer(sf.SendTimerTimer)

	sf.Timer = vcl.NewTimer(sf)
	sf.Timer.SetInterval(1000)
	sf.Timer.SetEnabled(true)
	sf.Timer.SetOnTimer(sf.TimerTimer)
}

func (sf *TMainMenu) Layout(sender *TFrmMain) {

	sf.MenuControl = &TMenuControl{
		TMenuItem: vcl.NewMenuItem(sf),
	}

	sf.MenuView = &TMenuView{
		TMenuItem: vcl.NewMenuItem(sf),
	}

	sf.MenuOption = &TMenuOption{
		TMenuItem: vcl.NewMenuItem(sf),
	}

	sf.N3 = &TMenuItem3{
		TMenuItem: vcl.NewMenuItem(sf),
	}

	sf.MenuControl.SetCaption("控制")
	sf.MenuControl.Layout(sender)

	sf.MenuView.SetCaption("查看")
	sf.MenuView.Layout(sender)

	sf.MenuOption.SetCaption("选项")
	sf.MenuOption.Layout(sender)

	sf.N3.SetCaption("帮助")
	sf.N3.Layout(sender)

	sf.Items().Add(sf.MenuControl)
	sf.Items().Add(sf.MenuView)
	sf.Items().Add(sf.MenuOption)
	sf.Items().Add(sf.N3)
}

func (sf *TPanel) Layout(sender *TFrmMain) {
	sf.SetParent(vcl.AsForm(sender))
	sf.SetAlign(types.AlTop)
	sf.SetBevelOuter(types.BvNone)
	sf.SetTabOrder(1)
	sf.SetLeft(0)
	sf.SetTop(0)
	sf.SetWidth(308)
	sf.SetHeight(119)

	sf.Label2 = vcl.NewLabel(sf)
	sf.Label2.SetParent(sf)
	sf.Label2.SetCaption("label2")
	sf.Label2.SetLeft(199)
	sf.Label2.SetTop(11)
	sf.Label2.SetWidth(42)
	sf.Label2.SetHeight(13)

	sf.Lack = vcl.NewLabel(sf)
	sf.Lack.SetParent(sf)
	sf.Lack.SetCaption("0/0")
	sf.Lack.SetLeft(195)
	sf.Lack.SetTop(33)
	sf.Lack.SetWidth(21)
	sf.Lack.SetHeight(13)

	sf.Hold = vcl.NewLabel(sf)
	sf.Hold.SetParent(sf)
	sf.Hold.SetCaption("")
	sf.Hold.SetLeft(106)
	sf.Hold.SetTop(10)
	sf.Hold.SetWidth(7)
	sf.Hold.SetHeight(13)
}

func (sf *TMenuControl) Layout(sender *TFrmMain) {
	sf.MenuControlStart = vcl.NewMenuItem(sf)
	sf.MenuControlStart.SetCaption("启动服务")
	sf.MenuControlStart.SetShortCutFromString("Ctrl+S")
	sf.MenuControlStart.SetOnClick(sender.MenuControlStartClick)

	sf.MenuControlStop = vcl.NewMenuItem(sf)
	sf.MenuControlStop.SetCaption("停止服务")
	sf.MenuControlStop.SetShortCutFromString("Ctrl+T")
	sf.MenuControlStop.SetOnClick(sender.MenuControlStopClick)

	sf.MenuControlReconnect = vcl.NewMenuItem(sf)
	sf.MenuControlReconnect.SetCaption("刷新连接")
	sf.MenuControlReconnect.SetShortCutFromString("Ctrl+R")
	sf.MenuControlReconnect.SetOnClick(sender.MenuControlReconnectClick)

	sf.MenuControlClearLog = vcl.NewMenuItem(sf)
	sf.MenuControlClearLog.SetCaption("清除日志")
	sf.MenuControlClearLog.SetShortCutFromString("Ctrl+C")
	sf.MenuControlClearLog.SetOnClick(sender.MenuControlClearLogClick)

	sf.MenuControlExit = vcl.NewMenuItem(sf)
	sf.MenuControlExit.SetCaption("退出")
	sf.MenuControlExit.SetShortCutFromString("Ctrl+X")
	sf.MenuControlExit.SetOnClick(sender.MenuControlExitClick)

	sf.Add(sf.MenuControlStart)
	sf.Add(sf.MenuControlStop)
	sf.Add(sf.MenuControlReconnect)
	sf.Add(sf.N1)
	sf.Add(sf.MenuControlClearLog)
	sf.Add(sf.N2)
	sf.Add(sf.MenuControlExit)
}

func (sf *TMenuView) Layout(sender *TFrmMain) {

	sf.MenuViewLogMsg = vcl.NewMenuItem(sf)
	sf.MenuViewLogMsg.SetCaption("查看日志")

	sf.Add(sf.MenuViewLogMsg)
}

func (sf *TMenuOption) Layout(sender *TFrmMain) {

	sf.MenuOptionGeneral = vcl.NewMenuItem(sf)
	sf.MenuOptionGeneral.SetCaption("基本设置")

	sf.MenuOptionIpFilter = vcl.NewMenuItem(sf)
	sf.MenuOptionIpFilter.SetCaption("安全过滤")
	sf.MenuOptionIpFilter.SetOnClick(sender.MenuOptionIpFilterClick)

	sf.Add(sf.MenuOptionGeneral)
	sf.Add(sf.MenuOptionIpFilter)
}

func (sf *TMenuItem3) Layout(sender *TFrmMain) {

	sf.N4 = vcl.NewMenuItem(sf)
	sf.N4.SetCaption("关于")
	sf.N4.SetOnClick(sender.N4Click)

	sf.Add(sf.N4)
}

// ******************** TFrmMain ********************
func (sf *TFrmMain) OnFormCreate(sender vcl.IObject) {

	sf.SetCaption("登陆网关")
	sf.EnabledMaximize(false)
	sf.SetBorderStyle(types.BsSingle)
	sf.SetLeft(636)
	sf.SetTop(215)
	sf.SetClientWidth(308)
	sf.SetClientHeight(154)
	// 布局
	sf.Layout()

	sf.tempLogList = make([]string, 0)
	sf.decodeMsgTime = 0
	sf.initUserSessionArray()
}

func (sf *TFrmMain) OnFormDestroy(sender vcl.IObject) {
	sf.tempLogList = sf.tempLogList[:0]
	for i := range SessionArray {
		SessionArray[i] = nil
	}
}

func (sf *TFrmMain) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
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
		userSession.RemoteIPaddr = ""
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

func (sf *TFrmMain) isBlockIP(ipaddr string) bool {
	ip := net.ParseIP(ipaddr)
	for i := range TempBlockIPList {
		ipAddr := TempBlockIPList[i]
		if ipAddr.IPaddr == int(ipToInteger(ip)) {
			return true
		}
	}
	for i := range BlockIPList {
		ipAddr := BlockIPList[i]
		if ipAddr.IPaddr == int(ipToInteger(ip)) {
			return true
		}
	}
	return false
}

func (sf *TFrmMain) isConnLimited(ipaddr string) bool {
	denyConnect := false
	ip := net.ParseIP(ipaddr)
	for i := range CurrIPaddrList {
		ipAddr := CurrIPaddrList[i]
		if ipAddr.IPaddr == int(ipToInteger(ip)) {
			ipAddr.Count++
			if (GetTickCount() - ipAddr.IPCountTick1) < 1000 {
				ipAddr.IPCount1++
				if ipAddr.IPCount1 >= IPCountLimit1 {
					denyConnect = true
				}
			} else {
				ipAddr.IPCountTick1 = GetTickCount()
				ipAddr.IPCount1 = 0
			}

			if (GetTickCount() - ipAddr.IPCountTick2) < 3000 {
				ipAddr.IPCount2++
				if ipAddr.IPCount2 >= IPCountLimit2 {
					denyConnect = true
				}
			} else {
				ipAddr.IPCountTick2 = GetTickCount()
				ipAddr.IPCount2 = 0
			}

			if ipAddr.Count > int(MaxConnOfIPaddr) {
				denyConnect = true
			}
		}
	}

	sockAddr := TSockaddr{
		IPaddr: int(ipToInteger(ip)),
		Count:  1,
	}
	CurrIPaddrList = append(CurrIPaddrList, sockAddr)
	return denyConnect
}

func (sf *TFrmMain) loadConfig() {
	conf := vcl.NewIniFile(ConfigFile)
	TitleName = conf.ReadString(GateClass, "Title", TitleName)
	ServerPort = conf.ReadInteger(GateClass, "ServerPort", ServerPort)
	ServerAddr = conf.ReadString(GateClass, "ServerAddr", ServerAddr)
	GatePort = conf.ReadInteger(GateClass, "GatePort", GatePort)
	GateAddr = conf.ReadString(GateClass, "GateAddr", GateAddr)
	ShowLogLevel = conf.ReadInteger(GateClass, "ShowLogLevel", ShowLogLevel)

	BlockMethod = TBlockIPMethod(conf.ReadInteger(GateClass, "BlockMethod", int32(BlockMethod)))

	if conf.ReadInteger(GateClass, "KeepConnectTimeOut", -1) <= 0 {
		conf.WriteInteger(GateClass, "KeepConnectTimeOut", KeepConnectTimeOut)
	}

	MaxConnOfIPaddr = conf.ReadInteger(GateClass, "MaxConnOfIPaddr", MaxConnOfIPaddr)
	KeepConnectTimeOut = conf.ReadInteger(GateClass, "KeepConnectTimeOut", KeepConnectTimeOut)
	DynamicIPDisMode = conf.ReadBool(GateClass, "DynamicIPDisMode", DynamicIPDisMode)

	conf.Free()

	LoadBlockIPFile()
}

func (sf *TFrmMain) resUserSessionArray() {
	for i := 0; i < GATEMAXSESSION; i++ {
		userSession := SessionArray[i]
		userSession.Socket = nil
		userSession.RemoteIPaddr = ""
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
	MainLogMsgListMutex.Lock()
	defer MainLogMsgListMutex.Unlock()

	if GetTickCount()-sf.showMainLogTick < 200 {
		return
	}
	sf.showMainLogTick = GetTickCount()

	sf.showLocked = true
	defer func() { sf.showLocked = false }()

	// 获取和清空主日志列表
	sf.tempLogList = append(sf.tempLogList, MainLogMsgList...)
	MainLogMsgList = MainLogMsgList[:0]

	// 在 GUI 中显示日志
	memoLog := vcl.AsMemo(sf.FindComponent("MemoLog"))
	for _, logMsg := range sf.tempLogList {
		vcl.ThreadSync(func() {
			memoLog.Lines().Add(logMsg)
		})
	}
	sf.tempLogList = sf.tempLogList[:0]
}

func (sf *TFrmMain) startService() {

	defer func() {
		if r := recover(); r != nil {
			sf.MainMenu.MenuControl.MenuControlStart.SetEnabled(true)
			sf.MainMenu.MenuControl.MenuControlStop.SetEnabled(false)
			MainOutMessage(fmt.Sprintf("%v", r), 0)
		}
	}()

	// 在这里添加启动服务的逻辑
	// 初始化变量和状态
	MainOutMessage("正在启动服务...", 3)
	ServiceStart = true
	GateReady = true
	sf.serverReady = false
	sf.sessionCount = 0
	sf.MainMenu.MenuControl.MenuControlStart.SetEnabled(false)
	sf.MainMenu.MenuControl.MenuControlStop.SetEnabled(true)

	sf.reConnectServerTick = GetTickCount() - 25*1000
	KeepAliveTimeOut = false
	SendMsgCount = 0
	TotalMsgListCount = 0
	sf.sendKeepAliveTick = GetTickCount()
	SendHoldTimeOut = false
	SendHoldTick = GetTickCount()

	CurrIPaddrList = make([]TSockaddr, 0)
	BlockIPList = make([]TSockaddr, 0)
	TempBlockIPList = make([]TSockaddr, 0)
	ClientSockeMsgList = make([]string, 0)

	sf.resUserSessionArray()
	sf.loadConfig()

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
	MainOutMessage("启动服务完成...", 3)
}

func (sf *TFrmMain) stopService() {

	MainOutMessage("正在停止服务...", 3)
	ServiceStart = false
	GateReady = false
	sf.MainMenu.MenuControl.MenuControlStart.SetEnabled(true)
	sf.MainMenu.MenuControl.MenuControlStop.SetEnabled(false)
	sf.SendTimer.SetEnabled(false)
	for i := 0; i < GATEMAXSESSION; i++ {
		if SessionArray[i].Socket != nil {
			SessionArray[i].Socket.Close()
		}
	}

	SaveBlockIPList()

	sf.ServerSocket.Close()
	sf.ClientSocket.Close()

	CurrIPaddrList = CurrIPaddrList[:0]
	BlockIPList = BlockIPList[:0]
	TempBlockIPList = TempBlockIPList[:0]
	ClientSockeMsgList = ClientSockeMsgList[:0]

	MainOutMessage("停止服务完成...", 3)
}

func (sf *TFrmMain) CloseConnect(ipaddr string) {
	if sf.ServerSocket.Active() {
		for {
			check := false
			for i := 0; i < sf.ServerSocket.ActiveConnections(); i++ {
				remoteAddr := getAddrHost(sf.ServerSocket.Connections()[i].RemoteAddr())
				if ipaddr == remoteAddr {
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
	sf.sessionCount = 0
	KeepAliveTick = GetTickCount()
	sf.resUserSessionArray()
	sf.serverReady = true
}

func (sf *TFrmMain) ClientSocketDisconnect(socket *TClientSocket, err error) {
	log.Info("ClientSocketDisconnect: {}", err.Error())
	for i := 0; i < GATEMAXSESSION; i++ {
		userSession := SessionArray[i]
		if userSession.Socket != nil {
			userSession.Socket.Close()
		}
		userSession.Socket = nil
		userSession.RemoteIPaddr = ""
		userSession.SocketHandle = uintptr(0)
	}

	sf.resUserSessionArray()
	ClientSockeMsgList = ClientSockeMsgList[:0]
	GateReady = false
	sf.sessionCount = 0
}

func (sf *TFrmMain) ClientSocketError(socket *TClientSocket, err error) {
	log.Info("ClientSocketError: {}", err.Error())
	socket.Close()
	sf.serverReady = false
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
			if common.TagCount(processMsg, '$') < 1 {
				break
			}
			processMsg, socketMsg = common.ArrestStringEx(processMsg, "%", "$")
			log.Info("processMsg: {}, socketMsg: {}", processMsg, socketMsg)
			if socketMsg == "" {
				break
			}
			if socketMsg[0] == '+' {
				if socketMsg[1] == '-' {
					sf.closeSocket(
						uintptr(
							common.StrToInt(socketMsg[2:len(socketMsg)-2], 0),
						),
					)
					continue
				} else {
					KeepAliveTick = GetTickCount()
					KeepAliveTimeOut = false
					continue
				}
			}
			socketMsg, socketHandleStr := common.GetValidStr3(socketMsg, []rune{'/'})
			socketHandle := uintptr(common.StrToInt(socketHandleStr, 0))
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
		if GetTickCount()-SessionArray[i].ConnctCheckTick > uint32(KeepConnectTimeOut) {
			remoteIPaddr := SessionArray[i].RemoteIPaddr
			switch BlockMethod {
			case Disconnect:
				SessionArray[i].Socket.Close()
			case Block:
				ip := net.ParseIP(remoteIPaddr)
				ipAddr := TSockaddr{}
				ipAddr.IPaddr = int(ipToInteger(ip))
				TempBlockIPList = append(TempBlockIPList, ipAddr)
				sf.CloseConnect(remoteIPaddr)
			case BlockList:
				ip := net.ParseIP(remoteIPaddr)
				ipAddr := TSockaddr{}
				ipAddr.IPaddr = int(ipToInteger(ip))
				BlockIPList = append(BlockIPList, ipAddr)
				sf.CloseConnect(remoteIPaddr)
			}
			MainOutMessage("端口空连接攻击: "+remoteIPaddr, 1)
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
					userSession.ConnctCheckTick = GetTickCount()
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
				MainOutMessage(userSession.IP+" : "+common.IntToStr(len(userSession.MsgList)), 5)
				SendMsgCount++
			} else {
				userSession.SocketHandle = uintptr(0)
				userSession.Socket = nil
				userSession.MsgList = userSession.MsgList[:0]
			}
		}
	}

	if GetTickCount()-sf.sendKeepAliveTick > 2*1000 {
		sf.sendKeepAliveTick = GetTickCount()
		if GateReady {
			sf.ClientSocket.Write([]byte("%--$"))
		}
	}

	if GetTickCount()-KeepAliveTick > 10*1000 {
		KeepAliveTimeOut = true
		sf.ClientSocket.Close()
	}

	decodeTime := GetTickCount() - decodeTick
	if sf.decodeMsgTime < decodeTime {
		sf.decodeMsgTime = decodeTime
	}
	if sf.decodeMsgTime > 50 {
		sf.decodeMsgTime -= 50
	}
}

func (sf *TFrmMain) MemoLogChange(sender vcl.IObject) {
	if sf.MemoLog.Lines().Count() > 200 {
		sf.MemoLog.Clear()
	}
}

func (sf *TFrmMain) MenuControlClearLogClick(sender vcl.IObject) {
	if vcl.MessageDlg("是否确认清除显示的日志信息?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.MrYes {
		vcl.ThreadSync(func() {
			sf.MemoLog.Clear()
		})
	}
}

func (sf *TFrmMain) MenuControlExitClick(sender vcl.IObject) {
	sf.Close()
}

func (sf *TFrmMain) MenuControlReconnectClick(sender vcl.IObject) {
	sf.reConnectServerTick = 0
}

func (sf *TFrmMain) MenuControlStartClick(sender vcl.IObject) {
	sf.startService()
}

func (sf *TFrmMain) MenuControlStopClick(sender vcl.IObject) {
	if vcl.MessageDlg("是否确认停止服务?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.MrYes {
		sf.stopService()
	}
}

func (sf *TFrmMain) MenuOptionGeneralClick(sender vcl.IObject) {
	//
}

func (sf *TFrmMain) MenuOptionIpFilterClick(sender vcl.IObject) {
	//
	FrmIPaddrFilter.ShowModal()
}

func (sf *TFrmMain) MenuViewLogMsgClick(sender vcl.IObject) {
	sf.MainMenu.MenuView.MenuViewLogMsg.SetChecked(!sf.MainMenu.MenuView.MenuViewLogMsg.Checked())
	sf.showLogMsg(sf.MainMenu.MenuView.MenuViewLogMsg.Checked())
}

func (sf *TFrmMain) N4Click(sender vcl.IObject) {
	MainLogOutMessage("引擎版本: 1.5.0 (20020522)")
	MainLogOutMessage("更新日期: 2023/09/14")
	MainLogOutMessage("程序制作: CHUNQIAN SHEN")
}

func (sf *TFrmMain) SendTimerTimer(sender vcl.IObject) {
	if sf.ServerSocket.Active() {
		ActiveConnections = sf.ServerSocket.ActiveConnections()
	}
	// 更新UI
	vcl.ThreadSync(func() {
		if SendHoldTimeOut {
			sf.Panel.Hold.SetCaption(common.IntToStr(ActiveConnections) + "#")
			if GetTickCount()-SendHoldTick > 3*1000 {
				SendHoldTimeOut = false
			}
		} else {
			sf.Panel.Hold.SetCaption(common.IntToStr(ActiveConnections))
		}
	})

	if GateReady && !KeepAliveTimeOut {
		for i := 0; i < GATEMAXSESSION; i++ {
			userSession := SessionArray[i]
			if userSession.Socket != nil {
				if GetTickCount()-userSession.UserTimeOutTick > 60*60*1000 {
					userSession.Socket.Close()
					userSession.Socket = nil
					userSession.SocketHandle = uintptr(0)
					userSession.MsgList = userSession.MsgList[:0]
					userSession.RemoteIPaddr = ""
				}
			}
		}
	}
	if !GateReady && ServiceStart {
		if GetTickCount()-sf.reConnectServerTick > 1000 /* 30*1000 */ {
			sf.reConnectServerTick = GetTickCount()
			sf.ClientSocket.Dial(sf, ServerAddr, ServerPort)
		}
	}
}

func (sf *TFrmMain) ServerSocketClientConnect(socket *TClientSocket) {
	var remoteIPaddr, localIPaddr string
	var ipAddr TSockaddr

	socket.Index = -1
	remoteIPaddr = getAddrHost(socket.RemoteAddr())

	if DynamicIPDisMode {
		localIPaddr = getAddrHost(sf.ClientSocket.RemoteAddr())
	} else {
		localIPaddr = getAddrHost(socket.LocalAddr())
	}

	if sf.isBlockIP(remoteIPaddr) {
		MainOutMessage("过滤连接: "+remoteIPaddr, 1)
		socket.Close()
		return
	}

	if sf.isConnLimited(remoteIPaddr) {
		switch BlockMethod {
		case Disconnect:
			socket.Close()
		case Block:
			ip := net.ParseIP(remoteIPaddr)
			ipAddr = TSockaddr{}
			ipAddr.IPaddr = int(ipToInteger(ip))
			TempBlockIPList = append(TempBlockIPList, ipAddr)
			sf.CloseConnect(remoteIPaddr)
		case BlockList:
			ip := net.ParseIP(remoteIPaddr)
			ipAddr = TSockaddr{}
			ipAddr.IPaddr = int(ipToInteger(ip))
			BlockIPList = append(BlockIPList, ipAddr)
			sf.CloseConnect(remoteIPaddr)
		}
		MainOutMessage("端口攻击: "+remoteIPaddr, 1)
	}

	if GateReady {
		for i := 0; i < GATEMAXSESSION; i++ {
			userSession := SessionArray[i]
			if userSession.Socket == nil {
				userSession.Socket = socket
				userSession.RemoteIPaddr = remoteIPaddr
				userSession.SendMsgLen = 0
				userSession.SendLock = false
				userSession.ConnctCheckTick = GetTickCount()
				userSession.SendAvailable = true
				userSession.SendCheck = false
				userSession.CheckSendLength = 0
				userSession.ReceiveLength = 0
				userSession.UserTimeOutTick = GetTickCount()
				userSession.SocketHandle = socket.SocketHandle()
				userSession.IP = remoteIPaddr
				userSession.MsgList = make([]string, 0)

				socket.Index = i
				sf.sessionCount++
				break
			}
		}

		if socket.Index >= 0 {
			// 发送字符串
			message := fmt.Sprintf("%%O%d/%s/%s$", socket.SocketHandle(), remoteIPaddr, localIPaddr)
			sf.ClientSocket.Write([]byte(message))
			MainOutMessage("Connect: "+remoteIPaddr, 5)
		} else {
			socket.Close()
			MainOutMessage("Kick Off: "+remoteIPaddr, 1)
		}
	} else {
		socket.Close()
		MainOutMessage("Kick Off: "+remoteIPaddr, 1)
	}
}

func (sf *TFrmMain) ServerSocketClientDisconnect(conn *TClientSocket, err error) {
	log.Info("ServerSocketClientDisconnect: {}", err.Error())
	remoteIPaddr := getAddrHost(conn.RemoteAddr())
	ip := net.ParseIP(remoteIPaddr)
	sockIndex := conn.Index

	for i := 0; i < len(CurrIPaddrList); i++ {
		ipAddr := CurrIPaddrList[i]
		if ipAddr.IPaddr == int(ipToInteger(ip)) {
			ipAddr.Count--
			if ipAddr.Count <= 0 {
				CurrIPaddrList = append(CurrIPaddrList[:i], CurrIPaddrList[i+1:]...)
			}
			break
		}
	}

	if sockIndex >= 0 && sockIndex < GATEMAXSESSION {
		userSession := SessionArray[sockIndex]
		userSession.Socket = nil
		userSession.RemoteIPaddr = ""
		userSession.SocketHandle = uintptr(0)
		userSession.MsgList = userSession.MsgList[:0]
		sf.sessionCount--
		if GateReady {
			message := fmt.Sprintf("%%X%d$", conn.SocketHandle())
			sf.ClientSocket.Write([]byte(message))
			MainOutMessage("DisConnect: "+remoteIPaddr, 5)
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
		if sf.serverReady {
			userSession.SendAvailable = true
			userSession.SendCheck = false
			userSession.CheckSendLength = 0
			if GateReady && !KeepAliveTimeOut {
				userSession.ConnctCheckTick = GetTickCount()
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
	startTimer := vcl.AsTimer(sender) // 将传入的 IObject 转型为 Timer
	if Started {
		startTimer.SetEnabled(false) // 禁用定时器
		sf.stopService()
		Close = true
		vcl.Application.Terminate() // 关闭应用程序
	} else {
		sf.MenuViewLogMsgClick(sender)
		Started = true
		startTimer.SetEnabled(false) // 禁用定时器
		sf.startService()
	}
}

func (sf *TFrmMain) TimerTimer(sender vcl.IObject) {
	var port string
	// 更新UI
	vcl.ThreadSync(func() {
		if sf.ServerSocket.Active() {
			port = getAddrPort(sf.ServerSocket.Addr())
			sf.StatusBar.Panels().Items(0).SetText(port)
			if SendHoldTimeOut {
				sf.StatusBar.Panels().Items(2).SetText(
					common.IntToStr(sf.sessionCount) + "/#" + common.IntToStr(sf.ServerSocket.ActiveConnections()),
				)
			} else {
				sf.StatusBar.Panels().Items(2).SetText(
					common.IntToStr(sf.sessionCount) + "/" + common.IntToStr(sf.ServerSocket.ActiveConnections()),
				)
			}
		} else {
			sf.StatusBar.Panels().Items(0).SetText("????")
			sf.StatusBar.Panels().Items(2).SetText("????")
		}
		sf.Panel.Label2.SetCaption(common.IntToStr(int(sf.decodeMsgTime)))
		if !GateReady {
			sf.StatusBar.Panels().Items(1).SetText("未连接")
		} else {
			if KeepAliveTimeOut {
				sf.StatusBar.Panels().Items(1).SetText("超时")
			} else {
				sf.StatusBar.Panels().Items(1).SetText("已连接")
				sf.Panel.Lack.SetCaption(common.IntToStr(TotalMsgListCount) + "/" + common.IntToStr(SendMsgCount))
			}
		}
	})
}

func MainLogOutMessage(msg string) {
	MainLogMsgListMutex.Lock()
	defer MainLogMsgListMutex.Unlock()

	MainLogMsgList = append(MainLogMsgList, msg)
}

func MainOutMessage(msg string, msgLevel int32) {
	MainLogMsgListMutex.Lock()
	defer MainLogMsgListMutex.Unlock()

	if msgLevel <= ShowLogLevel {
		tMsg := fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), msg)
		MainLogMsgList = append(MainLogMsgList, tMsg)
	}
}

func RGB(r, g, b uint32) types.TColor {
	return types.TColor(r | (g << 8) | (b << 16))
}
