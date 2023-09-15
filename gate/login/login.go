// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

// ******************** Const ********************
const (
	GATEMAXSESSION = 10000
)

// ******************** Type ********************
type TFrmMain struct {
	*vcl.TForm

	serverReady         bool
	showLocked          bool
	decodeMsgTime       uint32
	reConnectServerTick uint32
	sendKeepAliveTick   uint32
	showMainLogTick     uint32
	sessionCount        int
	tempLogList         []string
	tempLogListMutex    sync.Mutex

	ClientSocket         *net.TCPConn
	DecodeTimer          *vcl.TTimer
	Label2               *vcl.TLabel
	Hold                 *vcl.TLabel
	Lack                 *vcl.TLabel
	MainMenu             *vcl.TMainMenu
	MemoLog              *vcl.TMemo
	MenuControl          *vcl.TMenuItem
	MenuControlClearLog  *vcl.TMenuItem
	MenuControlExit      *vcl.TMenuItem
	MenuControlReconnect *vcl.TMenuItem
	MenuControlStart     *vcl.TMenuItem
	MenuControlStop      *vcl.TMenuItem
	MenuOption           *vcl.TMenuItem
	MenuOptionGeneral    *vcl.TMenuItem
	MenuOptionIpFilter   *vcl.TMenuItem
	MenuView             *vcl.TMenuItem
	MenuViewLogMsg       *vcl.TMenuItem
	N1                   *vcl.TMenuItem
	N2                   *vcl.TMenuItem
	N3                   *vcl.TMenuItem
	N4                   *vcl.TMenuItem
	Panel                *vcl.TPanel
	SendTimer            *vcl.TTimer
	ServerSocket         *net.TCPListener
	StartTimer           *vcl.TTimer
	StatusBar            *vcl.TStatusBar
	Timer                *vcl.TTimer
}

type TUserSession struct {
	Socket          *TServerSocket
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
	ConnctCheckTick uint32 // 连接数据传输空闲超时检测
}

type TSessionArray [GATEMAXSESSION]*TUserSession

// ******************** Var ********************
var (
	ClientSockeMsgList      []string
	ClientSockeMsgListMutex sync.Mutex
	FrmMain                 *TFrmMain
	SessionArray            TSessionArray
	ProcMsg                 string
)

// ******************** TFrmMain ********************
func (f *TFrmMain) OnFormCreate(sender vcl.IObject) {

	// ******************** TFrmMain ********************
	f.SetCaption("登陆网关")
	f.EnabledMaximize(false)
	f.SetBorderStyle(types.BsSingle)
	f.SetClientHeight(154)
	f.SetClientWidth(308)
	f.SetTop(215)
	f.SetLeft(636)

	// ******************** TMainMenu ********************
	f.MainMenu = vcl.NewMainMenu(f)

	f.MenuControl = vcl.NewMenuItem(f)
	f.MenuControlStart = vcl.NewMenuItem(f)
	f.MenuControlStop = vcl.NewMenuItem(f)
	f.MenuControlReconnect = vcl.NewMenuItem(f)
	f.MenuControlClearLog = vcl.NewMenuItem(f)
	f.MenuControlExit = vcl.NewMenuItem(f)

	f.MenuView = vcl.NewMenuItem(f)
	f.MenuViewLogMsg = vcl.NewMenuItem(f)

	f.MenuOption = vcl.NewMenuItem(f)
	f.MenuOptionGeneral = vcl.NewMenuItem(f)
	f.MenuOptionIpFilter = vcl.NewMenuItem(f)

	f.N1 = vcl.NewMenuItem(f)
	f.N2 = vcl.NewMenuItem(f)
	f.N3 = vcl.NewMenuItem(f)
	f.N4 = vcl.NewMenuItem(f)

	f.MenuControl.SetCaption("控制")
	f.MenuControlStart.SetCaption("启动服务")
	f.MenuControlStart.SetShortCutFromString("Ctrl+S")
	f.MenuControlStop.SetCaption("停止服务")
	f.MenuControlStop.SetShortCutFromString("Ctrl+T")
	f.MenuControlStop.SetOnClick(f.MenuControlStopClick)
	f.MenuControlReconnect.SetCaption("刷新连接")
	f.MenuControlReconnect.SetShortCutFromString("Ctrl+R")
	f.MenuControlClearLog.SetCaption("清除日志")
	f.MenuControlClearLog.SetShortCutFromString("Ctrl+C")
	f.MenuControlExit.SetCaption("退出")
	f.MenuControlExit.SetShortCutFromString("Ctrl+X")

	f.MenuView.SetCaption("查看")
	f.MenuViewLogMsg.SetCaption("查看日志")

	f.MenuOption.SetCaption("选项")
	f.MenuOptionGeneral.SetCaption("基本设置")
	f.MenuOptionIpFilter.SetCaption("安全过滤")

	f.N1.SetCaption("-")
	f.N2.SetCaption("-")

	f.N3.SetCaption("帮助")
	f.N4.SetCaption("关于")
	f.N4.SetOnClick(f.N4Click)

	f.MenuControl.Add(f.MenuControlStart)
	f.MenuControl.Add(f.MenuControlStop)
	f.MenuControl.Add(f.MenuControlReconnect)
	f.MenuControl.Add(f.N1)
	f.MenuControl.Add(f.MenuControlClearLog)
	f.MenuControl.Add(f.N2)
	f.MenuControl.Add(f.MenuControlExit)

	f.MenuView.Add(f.MenuViewLogMsg)

	f.MenuOption.Add(f.MenuOptionGeneral)
	f.MenuOption.Add(f.MenuOptionIpFilter)

	f.N3.Add(f.N4)

	f.MainMenu.Items().Add(f.MenuControl)
	f.MainMenu.Items().Add(f.MenuView)
	f.MainMenu.Items().Add(f.MenuOption)
	f.MainMenu.Items().Add(f.N3)

	// ******************** TMemo ********************
	f.MemoLog = vcl.NewMemo(f)
	f.MemoLog.SetName("MemoLog")
	f.MemoLog.SetAlign(types.AlClient)
	f.MemoLog.SetText("")
	f.MemoLog.SetParent(f)
	f.MemoLog.SetColor(colors.ClMenuText)
	f.MemoLog.Font().SetColor(colors.ClLimegreen)
	f.MemoLog.SetTop(119)
	f.MemoLog.SetLeft(0)
	f.MemoLog.SetHeight(18)
	f.MemoLog.SetWidth(308)
	f.MemoLog.SetWordWrap(false)
	f.MemoLog.SetScrollBars(types.SsHorizontal)

	// ******************** TPanel ********************
	f.Panel = vcl.NewPanel(f)
	f.Panel.SetParent(f)
	f.Panel.SetAlign(types.AlTop)
	f.Panel.SetBevelOuter(types.BvNone)
	f.Panel.SetTabOrder(1)
	f.Panel.SetTop(0)
	f.Panel.SetLeft(0)
	f.Panel.SetHeight(119)
	f.Panel.SetWidth(308)

	f.Label2 = vcl.NewLabel(f)
	f.Label2.SetParent(f.Panel)
	f.Label2.SetCaption("label2")
	f.Label2.SetTop(11)
	f.Label2.SetLeft(199)
	f.Label2.SetHeight(13)
	f.Label2.SetWidth(42)

	f.Lack = vcl.NewLabel(f)
	f.Lack.SetParent(f.Panel)
	f.Lack.SetCaption("0/0")
	f.Lack.SetTop(33)
	f.Lack.SetLeft(195)
	f.Lack.SetHeight(13)
	f.Lack.SetWidth(21)

	f.Hold = vcl.NewLabel(f)
	f.Hold.SetParent(f.Panel)
	f.Hold.SetCaption("")
	f.Hold.SetTop(10)
	f.Hold.SetLeft(106)
	f.Hold.SetHeight(13)
	f.Hold.SetWidth(7)

	// ******************** TStatusBar ********************
	f.StatusBar = vcl.NewStatusBar(f)
	f.StatusBar.SetParent(f)
	f.StatusBar.SetSimplePanel(false)
	f.StatusBar.SetTop(137)
	f.StatusBar.SetLeft(0)
	f.StatusBar.SetHeight(17)
	f.StatusBar.SetWidth(308)
	spnl := f.StatusBar.Panels().Add()
	spnl.SetAlignment(types.TaCenter)
	spnl.SetText("7100")
	spnl.SetWidth(50)
	spnl = f.StatusBar.Panels().Add()
	spnl.SetAlignment(types.TaCenter)
	spnl.SetText("未连接")
	spnl.SetWidth(60)
	spnl = f.StatusBar.Panels().Add()
	spnl.SetAlignment(types.TaCenter)
	spnl.SetText("0/0")
	spnl.SetWidth(70)
	spnl = f.StatusBar.Panels().Add()
	spnl.SetWidth(50)

	// ******************** TTimer ********************
	f.StartTimer = vcl.NewTimer(f)
	f.StartTimer.SetInterval(200)
	f.StartTimer.SetEnabled(true)
	f.StartTimer.SetOnTimer(f.StartTimerTimer)

	f.DecodeTimer = vcl.NewTimer(f)
	f.DecodeTimer.SetInterval(1)
	f.DecodeTimer.SetEnabled(true)
	f.DecodeTimer.SetOnTimer(f.DecodeTimerTimer)

	f.SendTimer = vcl.NewTimer(f)
	f.SendTimer.SetInterval(3000)
	f.SendTimer.SetEnabled(false)
	f.SendTimer.SetOnTimer(f.SendTimerTimer)

	// ******************** Non Visual ********************
	f.tempLogList = make([]string, 0)
	f.decodeMsgTime = 0
	f.initUserSessionArray()
}

func (f *TFrmMain) OnFormDestroy(sender vcl.IObject) {
	//
}

func (f *TFrmMain) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	*canClose = vcl.MessageDlg("是否确认退出服务器?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.IdYes
}

func (f *TFrmMain) closeSocket(socketHandle int) {
	//
}

func (f *TFrmMain) initUserSessionArray() {
	for i := 0; i < GATEMAXSESSION; i++ {
		session := TUserSession{}
		session.Socket = nil
		session.RemoteIPaddr = ""
		session.SendMsgLen = 0
		session.SendLock = false
		session.SendAvailable = true
		session.SendCheck = false
		session.CheckSendLength = 0
		session.ReceiveLength = 0
		session.UserTimeOutTick = GetTickCount()
		session.SocketHandle = uintptr(0)
		session.MsgList = make([]string, 0)
		SessionArray[i] = &session
	}
}

func (f *TFrmMain) isBlockIP(ipaddr string) bool {
	//
	return false
}

func (f *TFrmMain) isConnLimited(ipaddr string) bool {
	//
	return false
}

func (f *TFrmMain) loadConfig() {
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

func (f *TFrmMain) resUserSessionArray() {
	for i := 0; i < GATEMAXSESSION; i++ {
		userSession := SessionArray[i]
		userSession.Socket = nil
		userSession.RemoteIPaddr = ""
		userSession.SocketHandle = uintptr(0)
		userSession.MsgList = userSession.MsgList[:0]
	}
}

func (f *TFrmMain) sendUserMsg(userSession *TUserSession, sSendMsg string) int32 {
	return -1
}

func (f *TFrmMain) showLogMsg(boFlag bool) {
	var nHeight int32
	if boFlag {
		nHeight = f.Panel.Height()
		f.Panel.SetHeight(0)
		f.MemoLog.SetHeight(nHeight)
		f.MemoLog.SetTop(f.Panel.Top())
	} else {
		nHeight = f.MemoLog.Height()
		f.MemoLog.SetHeight(0)
		f.Panel.SetHeight(nHeight)
	}
}

func (f *TFrmMain) showMainLogMsg() {
	MainLogMsgListMutex.Lock()
	defer MainLogMsgListMutex.Unlock()

	if GetTickCount()-f.showMainLogTick < 200 {
		return
	}
	f.showMainLogTick = GetTickCount()

	f.showLocked = true
	defer func() { f.showLocked = false }()

	// 获取和清空主日志列表
	f.tempLogList = append(f.tempLogList, MainLogMsgList...)
	MainLogMsgList = MainLogMsgList[:0]

	// 在 GUI 中显示日志
	memoLog := vcl.AsMemo(f.FindComponent("MemoLog"))
	for _, logMsg := range f.tempLogList {
		vcl.ThreadSync(func() {
			memoLog.Lines().Add(logMsg)
		})
	}
	f.tempLogList = f.tempLogList[:0]
}

func (f *TFrmMain) startService() {

	defer func() {
		if r := recover(); r != nil {
			f.MenuControlStart.SetEnabled(true)
			f.MenuControlStop.SetEnabled(false)
			MainOutMessage(fmt.Sprintf("%v", r), 0)
		}
	}()

	// 在这里添加启动服务的逻辑
	// 初始化变量和状态
	MainOutMessage("正在启动服务...", 3)
	ServiceStart = true
	GateReady = true
	ServiceStart = true
	f.sessionCount = 0
	f.MenuControlStart.SetEnabled(false)
	f.MenuControlStop.SetEnabled(true)

	f.reConnectServerTick = GetTickCount() - 25*1000
	KeepAliveTimeOut = false
	SendMsgCount = 0
	TotalMsgListCount = 0
	f.sendKeepAliveTick = GetTickCount()
	SendHoldTimeOut = false
	SendHoldTick = GetTickCount()

	CurrIPaddrList = make([]TSockaddr, 0)
	BlockIPList = make([]TSockaddr, 0)
	TempBlockIPList = make([]TSockaddr, 0)
	ClientSockeMsgList = make([]string, 0)

	f.resUserSessionArray()
	f.loadConfig()

	if TitleName != "" {
		f.SetCaption(GateName + " - " + TitleName)
	} else {
		f.SetCaption(GateName)
	}

	// ClientSocket
	f.ClientSocketCreate()

	// ServerSocket
	f.ServerSocketCreate()

	f.SendTimer.SetEnabled(true)
	MainOutMessage("启动服务完成...", 3)
}

func (f *TFrmMain) stopService() {
	ClientSockeMsgListMutex.Lock()
	defer ClientSockeMsgListMutex.Unlock()

	MainOutMessage("正在停止服务...", 3)
	ServiceStart = false
	GateReady = false
	f.MenuControlStart.SetEnabled(true)
	f.MenuControlStop.SetEnabled(false)
	f.SendTimer.SetEnabled(false)
	for i := 0; i < GATEMAXSESSION; i++ {
		if SessionArray[i].Socket != nil {
			SessionArray[i].Socket.Close()
		}
	}

	SaveBlockIPList()

	f.ServerSocket.Close()
	f.ClientSocket.Close()

	CurrIPaddrList = CurrIPaddrList[:0]
	BlockIPList = BlockIPList[:0]
	TempBlockIPList = TempBlockIPList[:0]
	ClientSockeMsgList = ClientSockeMsgList[:0]

	MainOutMessage("停止服务完成...", 3)
}

func (f *TFrmMain) CloseConnect(ipaddr string) {
	//
}

func (f *TFrmMain) ClientSocketCreate() {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ServerAddr, ServerPort))
	if err != nil {
		panic("无法解析 Client 地址: " + err.Error())
	}
	f.ClientSocket, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic("无法监听 Client 地址: " + err.Error())
	}
}

func (f *TFrmMain) ClientSocketConnect(socket net.TCPConn) {
	//
}

func (f *TFrmMain) ClientSocketDisconnect(socket net.TCPConn) {
	//
}

func (f *TFrmMain) ClientSocketError(socket net.TCPConn, err error) {
	//
}

func (f *TFrmMain) ClientSocketRead(socket net.TCPConn) {
	//
}

func (f *TFrmMain) DecodeTimerTimer(sender vcl.IObject) {
	f.showMainLogMsg()
}

func (f *TFrmMain) MemoLogChange(sender vcl.IObject) {
	if f.MemoLog.Lines().Count() > 200 {
		f.MemoLog.Clear()
	}
}

func (f *TFrmMain) MenuControlClearLogClick(sender vcl.IObject) {
	//
}

func (f *TFrmMain) MenuControlExitClick(sender vcl.IObject) {
	//
}

func (f *TFrmMain) MenuControlReconnectClick(sender vcl.IObject) {
	//
}

func (f *TFrmMain) MenuControlStartClick(sender vcl.IObject) {
	//
}

func (f *TFrmMain) MenuControlStopClick(sender vcl.IObject) {
	if vcl.MessageDlg("是否确认停止服务?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.IdYes {
		f.stopService()
	}
}

func (f *TFrmMain) MenuOptionGeneralClick(sender vcl.IObject) {
	//
}

func (f *TFrmMain) MenuOptionIpFilterClick(sender vcl.IObject) {
	//
}

func (f *TFrmMain) MenuViewLogMsgClick(sender vcl.IObject) {
	f.MenuViewLogMsg.SetChecked(!f.MenuViewLogMsg.Checked())
	f.showLogMsg(f.MenuViewLogMsg.Checked())
}

func (f *TFrmMain) N4Click(sender vcl.IObject) {
	MainLogOutMessage("引擎版本: 1.5.0 (20020522)")
	MainLogOutMessage("更新日期: 2023/09/14")
	MainLogOutMessage("程序制作: CHUNQIAN SHEN")
}

func (f *TFrmMain) SendTimerTimer(sender vcl.IObject) {
	//
}

func (f *TFrmMain) ServerSocketCreate() {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", GateAddr, GatePort))
	if err != nil {
		panic("无法解析 Server 地址: " + err.Error())
	}
	f.ServerSocket, err = net.ListenTCP("tcp", addr)
	if err != nil {
		panic("无法监听 Server 地址: " + err.Error())
	}

	// 启动goroutine来接收Client Connect连接
}

func (f *TFrmMain) ServerSocketClientConnect(socket *TServerSocket) {
	var remoteIPaddr, localIPaddr string
	var ipAddr TSockaddr

	socket.Index = -1
	remoteIPaddr = socket.RemoteAddr().String()

	if DynamicIPDisMode {
		localIPaddr = f.ClientSocket.RemoteAddr().String()
	} else {
		localIPaddr = socket.LocalAddr().String()
	}

	if f.isBlockIP(remoteIPaddr) {
		MainOutMessage("过滤连接: "+remoteIPaddr, 1)
		socket.Close()
		return
	}

	if f.isConnLimited(remoteIPaddr) {
		switch BlockMethod {
		case Disconnect:
			socket.Close()
		case Block:
			ip := net.ParseIP(remoteIPaddr)
			ipAddr = TSockaddr{}
			ipAddr.IPaddr = int(ipToInteger(ip))
			TempBlockIPList = append(TempBlockIPList, ipAddr)
			f.CloseConnect(remoteIPaddr)
		case BlockList:
			ip := net.ParseIP(remoteIPaddr)
			ipAddr = TSockaddr{}
			ipAddr.IPaddr = int(ipToInteger(ip))
			BlockIPList = append(BlockIPList, ipAddr)
			f.CloseConnect(remoteIPaddr)
		}
		MainOutMessage("端口攻击: "+remoteIPaddr, 1)
	}

	if GateReady {
		for i := 0; i < GATEMAXSESSION; i++ {
			userSession := SessionArray[i]
			if userSession == nil {
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
				userSession.SocketHandle = socket.SocketHandle
				userSession.IP = remoteIPaddr
				userSession.MsgList = make([]string, 0)

				socket.Index = i
				f.sessionCount++
			}
		}

		if socket.Index >= 0 {
			// 发送字符串
			message := fmt.Sprintf("%%O%d/%s/%s$", socket.SocketHandle, remoteIPaddr, localIPaddr)
			f.ClientSocket.Write([]byte(message))
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

func (f *TFrmMain) ServerSocketClientDisconnect(conn *TServerSocket) {
	remoteIPaddr := conn.RemoteAddr().String()
	ip := net.ParseIP(remoteIPaddr)
	sockIndex := conn.Index

	for i := 0; i < len(CurrIPaddrList); i++ {
		ipAddr := CurrIPaddrList[i]
		if ipAddr.IPaddr == int(ipToInteger(ip)) {
			ipAddr.Count--
			if ipAddr.Count <= 0 {
				CurrIPaddrList = append(CurrIPaddrList[:i], CurrIPaddrList[i+1:]...)
			}
		}
	}

	if sockIndex >= 0 && sockIndex < GATEMAXSESSION {
		userSession := SessionArray[sockIndex]
		userSession.Socket = nil
		userSession.RemoteIPaddr = ""
		userSession.SocketHandle = uintptr(0)
		userSession.MsgList = userSession.MsgList[:0]
		f.sessionCount--
		if GateReady {
			message := fmt.Sprintf("%%X%d$", conn.SocketHandle)
			f.ClientSocket.Write([]byte(message))
			MainOutMessage("DisConnect: "+remoteIPaddr, 5)
		}
	}
}

func (f *TFrmMain) ServerSocketClientError(conn *TServerSocket, err error) {
	//
}

func (f *TFrmMain) ServerSocketClientRead(conn *TServerSocket) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	var dataBuffer bytes.Buffer
	reading := false

	sockIndex := conn.Index
	if sockIndex >= 0 && sockIndex < GATEMAXSESSION {
		userSession := SessionArray[sockIndex]
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				if err != io.EOF {
					fmt.Println("Read error:", err)
				}
				f.ServerSocketClientDisconnect(conn)
			}

			for i := 0; i < n; i++ {
				if buffer[i] == '%' {
					reading = true
					dataBuffer.Reset()
					continue
				}

				if buffer[i] == '$' {
					reading = false
					fmt.Println("Message Received:", dataBuffer.String())
					if f.serverReady {
						userSession.SendAvailable = true
						userSession.SendCheck = false
						userSession.CheckSendLength = 0
						if GateReady && !KeepAliveTimeOut {
							userSession.ConnctCheckTick = GetTickCount()
							if (GetTickCount() - userSession.UserTimeOutTick) < 1000 {
								userSession.ReceiveLength += dataBuffer.Len()
							} else {
								userSession.ReceiveLength = dataBuffer.Len()
							}
							message := fmt.Sprintf("%%A%d/%s$", conn.SocketHandle, dataBuffer.String())
							f.ClientSocket.Write([]byte(message))
						}
					}
					dataBuffer.Reset()
					continue
				}

				if reading {
					dataBuffer.WriteByte(buffer[i])
				}
			}
		}
	}
}

func (f *TFrmMain) StartTimerTimer(sender vcl.IObject) {
	startTimer := vcl.AsTimer(sender) // 将传入的 IObject 转型为 Timer
	if Started {
		startTimer.SetEnabled(false) // 禁用定时器
		f.stopService()
		Close = true
		vcl.Application.Terminate() // 关闭应用程序
	} else {
		f.MenuViewLogMsgClick(sender)
		Started = true
		startTimer.SetEnabled(false) // 禁用定时器
		f.startService()
	}
}

func (f *TFrmMain) TimerTimer(sender vcl.IObject) {
	//
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
