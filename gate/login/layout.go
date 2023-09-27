// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

func loginFormLayout(sender vcl.IObject) {
	f := &TFrmMain{
		TForm: vcl.AsForm(sender),
	}

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
	f.MenuControlStart.SetOnClick(f.MenuControlStartClick)

	f.MenuControlStop.SetCaption("停止服务")
	f.MenuControlStop.SetShortCutFromString("Ctrl+T")
	f.MenuControlStop.SetOnClick(f.MenuControlStopClick)

	f.MenuControlReconnect.SetCaption("刷新连接")
	f.MenuControlReconnect.SetShortCutFromString("Ctrl+R")
	f.MenuControlReconnect.SetOnClick(f.MenuControlReconnectClick)

	f.MenuControlClearLog.SetCaption("清除日志")
	f.MenuControlClearLog.SetShortCutFromString("Ctrl+C")
	f.MenuControlClearLog.SetOnClick(f.MenuControlClearLogClick)

	f.MenuControlExit.SetCaption("退出")
	f.MenuControlExit.SetShortCutFromString("Ctrl+X")
	f.MenuControlExit.SetOnClick(f.MenuControlExitClick)

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
	f.DecodeTimer.SetEnabled(false)
	f.DecodeTimer.SetOnTimer(f.DecodeTimerTimer)

	f.SendTimer = vcl.NewTimer(f)
	f.SendTimer.SetInterval(3000)
	f.SendTimer.SetEnabled(false)
	f.SendTimer.SetOnTimer(f.SendTimerTimer)

	f.Timer = vcl.NewTimer(f)
	f.Timer.SetInterval(1000)
	f.Timer.SetEnabled(true)
	f.Timer.SetOnTimer(f.TimerTimer)
}

func loginFilterLayout(sender vcl.IObject) {
}

func loginConfigLayout(sender vcl.IObject) {
}
