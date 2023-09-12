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
)

// ******************** Const ********************
const (
	GATEMAXSESSION = 10000
)

// ******************** Type ********************
type TFrmMain struct {
	*vcl.TForm

	boServerReady         bool
	boShowLocked          bool
	dwDecodeMsgTime       uint32
	dwReConnectServerTick uint32
	dwSendKeepAliveTick   uint32
	dwShowMainLogTick     uint32
	nSessionCount         int
	stringList30C         []string
	stringList318         []string
	tempLogList           []string
	tempLogListMutex      sync.Mutex

	ClientSocket         *net.TCPConn
	DecodeTimer          *vcl.TTimer
	Label2               *vcl.TLabel
	LbHold               *vcl.TLabel
	LbLack               *vcl.TLabel
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
	ServerSocket         *net.TCPConn
	StartTimer           *vcl.TTimer
	StatusBar            *vcl.TStatusBar
	Timer                *vcl.TTimer
}

type TUserSession struct {
	Socket            *net.Conn // 0x00
	SRemoteIPaddr     string    // 0x04
	NSendMsgLen       int       // 0x08
	Bo0C              bool      // 0x0C
	Dw10Tick          uint32    // 0x10
	NCheckSendLength  int       // 0x14
	BoSendAvailable   bool      // 0x18
	BoSendCheck       bool      // 0x19
	DwSendLockTimeOut uint32    // 0x1C
	N20               int       // 0x20
	DwUserTimeOutTick uint32    // 0x24
	SocketHandle      int       // 0x28
	SIP               string    // 0x2C
	MsgList           []string  // 0x30
	MsgListMutex      sync.Mutex
	DwConnctCheckTick uint32 // 连接数据传输空闲超时检测
}

type TSessionArray [GATEMAXSESSION]TUserSession

// ******************** Var ********************
var (
	ClientSockeMsgList      []string
	FrmMain                 *TFrmMain
	GSessionArray           *TSessionArray
	SProcMsg                string
	ClientSockeMsgListMutex sync.Mutex
	GSessionArrayMutex      sync.Mutex
)

// ******************** TFrmMain ********************
func (f *TFrmMain) closeSocket(nSocketHandle int) {
	//
}

func (f *TFrmMain) iniUserSessionArray() {
	GSessionArrayMutex.Lock()
	defer GSessionArrayMutex.Unlock()

	for i := 0; i < GATEMAXSESSION; i++ {
		session := &GSessionArray[i]
		session.Socket = nil
		session.SRemoteIPaddr = ""
		session.NSendMsgLen = 0
		session.Bo0C = false
		session.Dw10Tick = GetTickCount()
		session.NCheckSendLength = 0
		session.BoSendAvailable = true
		session.BoSendCheck = false
		session.N20 = 0
		session.DwUserTimeOutTick = GetTickCount()
		session.SocketHandle = -1
		session.MsgList = make([]string, 0)
	}
}

func (f *TFrmMain) isBlockIP(sIPaddr string) bool {
	//
	return false
}

func (f *TFrmMain) isConnLimited(sIPaddr string) bool {
	//
	return false
}

func (f *TFrmMain) loadConfig() {
	//
}

func (f *TFrmMain) resUserSessionArray() {
	//
}

func (f *TFrmMain) sendUserMsg(userSession *TUserSession, sSendMsg string) int {
	//
	return 0
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

	if GetTickCount()-f.dwShowMainLogTick < 200 {
		return
	}
	f.dwShowMainLogTick = GetTickCount()

	f.boShowLocked = true
	defer func() { f.boShowLocked = false }()

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
	// 在这里添加启动服务的逻辑

	// 初始化变量和状态
	MainOutMessage("正在启动服务...", 3)
	BoServiceStart = true
	BoGateReady = true
	BoServiceStart = true
	f.nSessionCount = 0
	f.MenuControlStart.SetEnabled(false)
	f.MenuControlStop.SetEnabled(true)

	f.dwReConnectServerTick = GetTickCount() - 25*1000
	BoKeepAliveTimeOut = false
	NSendMsgCount = 0
	N456A2C = 0
	f.dwSendKeepAliveTick = GetTickCount()
	BoSendHoldTimeOut = false
	DwSendHoldTick = GetTickCount()

	CurrIPaddrList = make([]TSockaddr, 0)
	BlockIPList = make([]TSockaddr, 0)
	TempBlockIPList = make([]TSockaddr, 0)
	ClientSockeMsgList = make([]string, 0)

	// ResUserSessionArray()
	// LoadConfig()

	if TitleName != "" {
		f.SetCaption(GateName + " - " + TitleName)
	} else {
		f.SetCaption(GateName)
	}

	// f.SendTimer.SetEnabled(true)
	MainOutMessage("启动服务完成...", 3)
}

func (f *TFrmMain) stopService() {
	//
}

func (f *TFrmMain) CloseConnect(sIPaddr string) {
	//
}

func (f *TFrmMain) ClientSocketConnect(sender vcl.IObject, socket net.TCPConn) {
	//
}

func (f *TFrmMain) ClientSocketDisconnect(sender vcl.IObject, socket net.TCPConn) {
	//
}

func (f *TFrmMain) ClientSocketError(sender vcl.IObject, socket net.TCPConn, errorEvent error) {
	//
}

func (f *TFrmMain) ClientSocketRead(sender vcl.IObject, socket net.TCPConn) {
	//
}

func (f *TFrmMain) DecodeTimerTimer(sender vcl.IObject) {
	f.showMainLogMsg()
}

func (f *TFrmMain) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	*canClose = vcl.MessageDlg("是否确认退出服务器?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.IdYes
}

func (f *TFrmMain) OnFormCreate(sender vcl.IObject) {
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

	f.LbLack = vcl.NewLabel(f)
	f.LbLack.SetParent(f.Panel)
	f.LbLack.SetCaption("0/0")
	f.LbLack.SetTop(33)
	f.LbLack.SetLeft(195)
	f.LbLack.SetHeight(13)
	f.LbLack.SetWidth(21)

	f.LbHold = vcl.NewLabel(f)
	f.LbHold.SetParent(f.Panel)
	f.LbHold.SetCaption("")
	f.LbHold.SetTop(10)
	f.LbHold.SetLeft(106)
	f.LbHold.SetHeight(13)
	f.LbHold.SetWidth(7)

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
}

func (f *TFrmMain) OnFormDestroy(sender vcl.IObject) {
	//
}

func (f *TFrmMain) MemoLogChange(sender vcl.IObject) {
	//
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
	//
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
	//
}

func (f *TFrmMain) SendTimerTimer(sender vcl.IObject) {
	//
}

func (f *TFrmMain) ServerSocketClientConnect(sender vcl.IObject, socket net.TCPConn) {
	//
}

func (f *TFrmMain) ServerSocketClientDisconnect(sender vcl.IObject, socket net.TCPConn) {
	//
}

func (f *TFrmMain) ServerSocketClientError(sender vcl.IObject, socket net.TCPConn, errorEvent error) {
	//
}

func (f *TFrmMain) ServerSocketClientRead(sender vcl.IObject, socket net.TCPConn) {
	//
}

func (f *TFrmMain) StartTimerTimer(sender vcl.IObject) {
	startTimer := vcl.AsTimer(sender) // 将传入的 IObject 转型为 Timer
	if BoStarted {
		startTimer.SetEnabled(false) // 禁用定时器
		f.stopService()
		BoClose = true
		vcl.Application.Terminate() // 关闭应用程序
	} else {
		f.MenuViewLogMsgClick(sender)
		BoStarted = true
		startTimer.SetEnabled(false) // 禁用定时器
		f.startService()
	}
}

func (f *TFrmMain) TimerTimer(sender vcl.IObject) {
	//
}

func RGB(r, g, b uint32) types.TColor {
	return types.TColor(r | (g << 8) | (b << 16))
}

func MainLogOutMessage(sMsg string) {
	//
}

func MainOutMessage(sMsg string, nMsgLevel int) {
	MainLogMsgListMutex.Lock()
	defer MainLogMsgListMutex.Unlock()

	if nMsgLevel <= NShowLogLevel {
		tMsg := fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), sMsg)
		MainLogMsgList = append(MainLogMsgList, tMsg)
	}
}
