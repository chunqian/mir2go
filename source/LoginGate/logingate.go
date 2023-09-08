// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import "github.com/ying32/govcl/vcl"

type TFrmMain struct {
	*vcl.TForm

    memoLog *vcl.TMemo
    panel *vcl.TPanel
    statusBar *vcl.TStatusBar
    mainMenu *vcl.TMainMenu

    lbHold *vcl.TLabel
    lbLack *vcl.TLabel
    label2 *vcl.TLabel
    
    startTimer *vcl.TTimer
    sendTimer *vcl.TTimer
    decodeTimer *vcl.TTimer

    menuControl *vcl.TMenuItem
    menuControlStart *vcl.TMenuItem
    menuControlStop *vcl.TMenuItem
    menuControlReconnect *vcl.TMenuItem
    menuControlClearLog *vcl.TMenuItem
    menuControlExit *vcl.TMenuItem

    menuView *vcl.TMenuItem
    menuViewLogMsg *vcl.TMenuItem

    menuOption *vcl.TMenuItem
    menuOptionGeneral *vcl.TMenuItem
    menuOptionIpFilter *vcl.TMenuItem

    n1 *vcl.TMenuItem
    n2 *vcl.TMenuItem
    n3 *vcl.TMenuItem
    n4 *vcl.TMenuItem

    // clientSocket *vcl.TClientSocket
    // serverSocket *vcl.TServerSocket
}

var (
	frmMain *TFrmMain
)

// ******************** TFrmMain ********************
func (f *TFrmMain) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("登陆网关")
	f.EnabledMaximize(false)
	f.SetClientHeight(154)
	f.SetClientWidth(308)
	f.SetTop(215)
	f.SetLeft(636)

	f.mainMenu = vcl.NewMainMenu(f)

	f.menuControl = vcl.NewMenuItem(f)
	f.menuControlStart = vcl.NewMenuItem(f)
	f.menuControlStop = vcl.NewMenuItem(f)
	f.menuControlReconnect = vcl.NewMenuItem(f)
	f.menuControlClearLog = vcl.NewMenuItem(f)
	f.menuControlExit = vcl.NewMenuItem(f)

	f.n1 = vcl.NewMenuItem(f)
	f.n2 = vcl.NewMenuItem(f)

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

	f.n1.SetCaption("-")
	f.n2.SetCaption("-")

	f.menuControl.Add(f.menuControlStart)
	f.menuControl.Add(f.menuControlStop)
	f.menuControl.Add(f.menuControlReconnect)
	f.menuControl.Add(f.n1)
	f.menuControl.Add(f.menuControlClearLog)
	f.menuControl.Add(f.n2)
	f.menuControl.Add(f.menuControlExit)

	f.mainMenu.Items().Add(f.menuControl)
}