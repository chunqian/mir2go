// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"net"
	"sync"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"

	"github.com/chunqian/mir2go/common"
)

type TFrmMain struct {
	*vcl.TForm

	memoLog   *vcl.TMemo
	panel     *vcl.TPanel
	statusBar *vcl.TStatusBar
	mainMenu  *vcl.TMainMenu

	lbHold *vcl.TLabel
	lbLack *vcl.TLabel
	label2 *vcl.TLabel

	startTimer  *vcl.TTimer
	sendTimer   *vcl.TTimer
	decodeTimer *vcl.TTimer

	menuControl          *vcl.TMenuItem
	menuControlStart     *vcl.TMenuItem
	menuControlStop      *vcl.TMenuItem
	menuControlReconnect *vcl.TMenuItem
	menuControlClearLog  *vcl.TMenuItem
	menuControlExit      *vcl.TMenuItem

	menuView       *vcl.TMenuItem
	menuViewLogMsg *vcl.TMenuItem

	menuOption         *vcl.TMenuItem
	menuOptionGeneral  *vcl.TMenuItem
	menuOptionIpFilter *vcl.TMenuItem

	n1 *vcl.TMenuItem
	n2 *vcl.TMenuItem
	n3 *vcl.TMenuItem
	n4 *vcl.TMenuItem

	clientSocket *net.TCPConn
	serverSocket *net.TCPConn

	dwShowMainLogTick     uint32     // LongWord, 显示主日志的时间戳
	boShowLocked          bool       // Boolean, 是否锁定显示
	tempLogList           []string   // 暂存日志列表
	nSessionCount         int        // Integer, 会话数量
	stringList30C         []string   // 特定的字符串列表
	dwSendKeepAliveTick   uint32     // LongWord, 发送心跳包的时间戳
	boServerReady         bool       // Boolean, 服务器是否准备好
	stringList318         []string   // 特定的字符串列表
	dwDecodeMsgTime       uint32     // LongWord, 解码消息所需时间
	dwReConnectServerTick uint32     // LongWord, 重新连接服务器的时间戳
	mutex_stringList      sync.Mutex // 用于并发场景下安全访问TStringList
}

var (
	frmMain *TFrmMain

	g_SessionArray     *common.TSessionArray
	mutex_SessionArray sync.Mutex

	clientSockeMsgList *vcl.TStringList
	sProcMsg           string
)

// ******************** TFrmMain ********************
func (f *TFrmMain) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("登陆网关")
	f.EnabledMaximize(false)
	f.SetBorderStyle(types.BsSingle)
	f.SetClientHeight(154)
	f.SetClientWidth(308)
	f.SetTop(215)
	f.SetLeft(636)

	// ******************** TMainMenu ********************
	f.mainMenu = vcl.NewMainMenu(f)

	f.menuControl = vcl.NewMenuItem(f)
	f.menuControlStart = vcl.NewMenuItem(f)
	f.menuControlStop = vcl.NewMenuItem(f)
	f.menuControlReconnect = vcl.NewMenuItem(f)
	f.menuControlClearLog = vcl.NewMenuItem(f)
	f.menuControlExit = vcl.NewMenuItem(f)

	f.menuView = vcl.NewMenuItem(f)
	f.menuViewLogMsg = vcl.NewMenuItem(f)

	f.menuOption = vcl.NewMenuItem(f)
	f.menuOptionGeneral = vcl.NewMenuItem(f)
	f.menuOptionIpFilter = vcl.NewMenuItem(f)

	f.n1 = vcl.NewMenuItem(f)
	f.n2 = vcl.NewMenuItem(f)
	f.n3 = vcl.NewMenuItem(f)
	f.n4 = vcl.NewMenuItem(f)

	f.menuControl.SetCaption("控制")
	f.menuControlStart.SetCaption("启动服务")
	f.menuControlStart.SetShortCutFromString("Ctrl+S")
	f.menuControlStop.SetCaption("停止服务")
	f.menuControlStop.SetShortCutFromString("Ctrl+T")
	f.menuControlReconnect.SetCaption("刷新连接")
	f.menuControlReconnect.SetShortCutFromString("Ctrl+R")
	f.menuControlClearLog.SetCaption("清除日志")
	f.menuControlClearLog.SetShortCutFromString("Ctrl+C")
	f.menuControlExit.SetCaption("退出")
	f.menuControlExit.SetShortCutFromString("Ctrl+X")

	f.menuView.SetCaption("查看")
	f.menuViewLogMsg.SetCaption("查看日志")

	f.menuOption.SetCaption("选项")
	f.menuOptionGeneral.SetCaption("基本设置")
	f.menuOptionIpFilter.SetCaption("安全过滤")

	f.n1.SetCaption("-")
	f.n2.SetCaption("-")

	f.n3.SetCaption("帮助")
	f.n4.SetCaption("关于")

	f.menuControl.Add(f.menuControlStart)
	f.menuControl.Add(f.menuControlStop)
	f.menuControl.Add(f.menuControlReconnect)
	f.menuControl.Add(f.n1)
	f.menuControl.Add(f.menuControlClearLog)
	f.menuControl.Add(f.n2)
	f.menuControl.Add(f.menuControlExit)

	f.menuView.Add(f.menuViewLogMsg)

	f.menuOption.Add(f.menuOptionGeneral)
	f.menuOption.Add(f.menuOptionIpFilter)

	f.n3.Add(f.n4)

	f.mainMenu.Items().Add(f.menuControl)
	f.mainMenu.Items().Add(f.menuView)
	f.mainMenu.Items().Add(f.menuOption)
	f.mainMenu.Items().Add(f.n3)

	// ******************** TMemo ********************
	f.memoLog = vcl.NewMemo(f)
	f.memoLog.SetName("MemoLog")
	f.memoLog.SetAlign(types.AlClient)
	f.memoLog.SetText("")
	f.memoLog.SetParent(f)
	f.memoLog.SetColor(colors.ClMenuText)
	f.memoLog.Font().SetColor(colors.ClLimegreen)
	f.memoLog.SetTop(119)
	f.memoLog.SetLeft(0)
	f.memoLog.SetHeight(18)
	f.memoLog.SetWidth(308)
	f.memoLog.SetWordWrap(false)
	f.memoLog.SetScrollBars(types.SsHorizontal)

	// ******************** TPanel ********************
	f.panel = vcl.NewPanel(f)
	f.panel.SetParent(f)
	f.panel.SetAlign(types.AlTop)
	f.panel.SetBevelOuter(types.BvNone)
	f.panel.SetTabOrder(1)
	f.panel.SetTop(0)
	f.panel.SetLeft(0)
	f.panel.SetHeight(119)
	f.panel.SetWidth(308)

	f.label2 = vcl.NewLabel(f)
	f.label2.SetParent(f.panel)
	f.label2.SetCaption("label2")
	f.label2.SetTop(11)
	f.label2.SetLeft(199)
	f.label2.SetHeight(13)
	f.label2.SetWidth(42)

	f.lbLack = vcl.NewLabel(f)
	f.lbLack.SetParent(f.panel)
	f.lbLack.SetCaption("0/0")
	f.lbLack.SetTop(33)
	f.lbLack.SetLeft(195)
	f.lbLack.SetHeight(13)
	f.lbLack.SetWidth(21)

	f.lbHold = vcl.NewLabel(f)
	f.lbHold.SetParent(f.panel)
	f.lbHold.SetCaption("")
	f.lbHold.SetTop(10)
	f.lbHold.SetLeft(106)
	f.lbHold.SetHeight(13)
	f.lbHold.SetWidth(7)

	// ******************** TStatusBar ********************
	f.statusBar = vcl.NewStatusBar(f)
	f.statusBar.SetParent(f)
	f.statusBar.SetSimplePanel(false)
	f.statusBar.SetTop(137)
	f.statusBar.SetLeft(0)
	f.statusBar.SetHeight(17)
	f.statusBar.SetWidth(308)
	spnl := f.statusBar.Panels().Add()
	spnl.SetAlignment(types.TaCenter)
	spnl.SetText("7100")
	spnl.SetWidth(50)
	spnl = f.statusBar.Panels().Add()
	spnl.SetAlignment(types.TaCenter)
	spnl.SetText("未连接")
	spnl.SetWidth(60)
	spnl = f.statusBar.Panels().Add()
	spnl.SetAlignment(types.TaCenter)
	spnl.SetText("0/0")
	spnl.SetWidth(70)
	spnl = f.statusBar.Panels().Add()
	spnl.SetWidth(50)

	// ******************** TTimer ********************

	f.startTimer = vcl.NewTimer(f)
	f.startTimer.SetInterval(200)
	f.startTimer.SetEnabled(true)
	f.startTimer.SetOnTimer(f.StartTimerTimer)

	f.decodeTimer = vcl.NewTimer(f)
	f.decodeTimer.SetInterval(1)
	f.decodeTimer.SetEnabled(true)
	f.decodeTimer.SetOnTimer(f.DecodeTimerTimer)
}

func (f *TFrmMain) OnFormCloseQuery(Sender vcl.IObject, CanClose *bool) {
	*CanClose = vcl.MessageDlg("是否确认退出服务器?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.IdYes
}

func RGB(r, g, b uint32) types.TColor {
	return types.TColor(r | (g << 8) | (b << 16))
}

func (f *TFrmMain) sendUserMsg(UserSession *common.TSessionArray, sSendMsg string) int {
	return 0
}

func (f *TFrmMain) IniUserSessionArray() {
	mutex_SessionArray.Lock()
	defer mutex_SessionArray.Unlock()

	for i := 0; i < common.GATEMAXSESSION; i++ {
		session := &g_SessionArray[i]
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

func (f *TFrmMain) StartTimerTimer(sender vcl.IObject) {
	startTimer := vcl.AsTimer(sender) // 将传入的 IObject 转型为 Timer
	if boStarted {
		startTimer.SetEnabled(false) // 禁用定时器
		f.StopService()
		boClose = true
		vcl.Application.Terminate() // 关闭应用程序
	} else {
		f.MenuViewLogMsgClick(sender)
		boStarted = true
		startTimer.SetEnabled(false) // 禁用定时器
		f.StartService()
	}
}

func (f *TFrmMain) DecodeTimerTimer(sender vcl.IObject) {
	ShowMainLogMsg(f)
}

func (f *TFrmMain) StopService() {
	// 在这里添加关闭服务的逻辑
}

func (f *TFrmMain) StartService() {
	// 在这里添加启动服务的逻辑

	// 初始化变量和状态
	MainOutMessage("正在启动服务...", 3)
}

func (f *TFrmMain) MenuViewLogMsgClick(sender vcl.IObject) {
	// 在这里添加与菜单项点击相关的逻辑
	f.menuViewLogMsg.SetChecked(!f.menuViewLogMsg.Checked())
	f.ShowLogMsg(f.menuViewLogMsg.Checked())
}

// ShowLogMsg 对应于 Delphi 中的 ShowLogMsg
func (f *TFrmMain) ShowLogMsg(boFlag bool) {
	var nHeight int32
	if boFlag {
		nHeight = f.panel.Height()
		f.panel.SetHeight(0)
		f.memoLog.SetHeight(nHeight)
		f.memoLog.SetTop(f.panel.Top())
	} else {
		nHeight = f.memoLog.Height()
		f.memoLog.SetHeight(0)
		f.panel.SetHeight(nHeight)
	}
}
