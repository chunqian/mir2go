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

	dwShowMainLogTick     uint32           // LongWord, 显示主日志的时间戳
	boShowLocked          bool             // Boolean, 是否锁定显示
	tempLogList           *vcl.TStringList // 暂存日志列表
	nSessionCount         int              // Integer, 会话数量
	stringList30C         *vcl.TStringList // 特定的字符串列表
	dwSendKeepAliveTick   uint32           // LongWord, 发送心跳包的时间戳
	boServerReady         bool             // Boolean, 服务器是否准备好
	stringList318         *vcl.TStringList // 特定的字符串列表
	dwDecodeMsgTime       uint32           // LongWord, 解码消息所需时间
	dwReConnectServerTick uint32           // LongWord, 重新连接服务器的时间戳
	mutex                 sync.Mutex       // 用于并发场景下安全访问TStringList
}

var (
	frmMain *TFrmMain
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
	f.memoLog.SetParent(f)
	f.memoLog.SetColor(colors.ClMenuText)
	f.memoLog.SetTop(119)
	f.memoLog.SetLeft(0)
	f.memoLog.SetHeight(18)
	f.memoLog.SetWidth(308)
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

	// ******************** Other ********************
}

func (f *TFrmMain) OnFormCloseQuery(Sender vcl.IObject, CanClose *bool) {
	*CanClose = vcl.MessageDlg("是否确认退出服务器?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.IdYes
}

func (f *TFrmMain) sendUserMsg(UserSession *common.TSessionArray, sSendMsg string) int {
	return 0
}
