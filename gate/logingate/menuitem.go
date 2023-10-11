// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	. "github.com/chunqian/mir2go/common"
)

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

func NewMenuControl(sender vcl.IComponent) *TMenuControl {
	sf := &TMenuControl{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.MenuControlStart = vcl.NewMenuItem(sf)
	sf.MenuControlStart.SetCaption("启动服务")
	sf.MenuControlStart.SetShortCutFromString("Ctrl+S")
	sf.MenuControlStart.SetOnClick(sf.MenuControlStartClick)

	sf.MenuControlStop = vcl.NewMenuItem(sf)
	sf.MenuControlStop.SetCaption("停止服务")
	sf.MenuControlStop.SetShortCutFromString("Ctrl+T")
	sf.MenuControlStop.SetOnClick(sf.MenuControlStopClick)

	sf.MenuControlReconnect = vcl.NewMenuItem(sf)
	sf.MenuControlReconnect.SetCaption("刷新连接")
	sf.MenuControlReconnect.SetShortCutFromString("Ctrl+R")
	sf.MenuControlReconnect.SetOnClick(sf.MenuControlReconnectClick)

	sf.MenuControlClearLog = vcl.NewMenuItem(sf)
	sf.MenuControlClearLog.SetCaption("清除日志")
	sf.MenuControlClearLog.SetShortCutFromString("Ctrl+C")
	sf.MenuControlClearLog.SetOnClick(sf.MenuControlClearLogClick)

	sf.MenuControlExit = vcl.NewMenuItem(sf)
	sf.MenuControlExit.SetCaption("退出")
	sf.MenuControlExit.SetShortCutFromString("Ctrl+X")
	sf.MenuControlExit.SetOnClick(sf.MenuControlExitClick)

	sf.Add(sf.MenuControlStart)
	sf.Add(sf.MenuControlStop)
	sf.Add(sf.MenuControlReconnect)
	sf.Add(sf.N1)
	sf.Add(sf.MenuControlClearLog)
	sf.Add(sf.N2)
	sf.Add(sf.MenuControlExit)

	return sf
}

func NewMenuView(sender vcl.IComponent) *TMenuView {
	sf := &TMenuView{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.MenuViewLogMsg = vcl.NewMenuItem(sf)
	sf.MenuViewLogMsg.SetCaption("查看日志")
	sf.MenuViewLogMsg.SetOnClick(sf.MenuViewLogMsgClick)

	sf.Add(sf.MenuViewLogMsg)

	return sf
}

func NewMenuOption(sender vcl.IComponent) *TMenuOption {
	sf := &TMenuOption{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.MenuOptionGeneral = vcl.NewMenuItem(sf)
	sf.MenuOptionGeneral.SetCaption("基本设置")
	sf.MenuOptionGeneral.SetOnClick(sf.MenuOptionGeneralClick)

	sf.MenuOptionIpFilter = vcl.NewMenuItem(sf)
	sf.MenuOptionIpFilter.SetCaption("安全过滤")
	sf.MenuOptionIpFilter.SetOnClick(sf.MenuOptionIpFilterClick)

	sf.Add(sf.MenuOptionGeneral)
	sf.Add(sf.MenuOptionIpFilter)

	return sf
}

func NewMenuItem3(sender vcl.IComponent) *TMenuItem3 {
	sf := &TMenuItem3{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.N4 = vcl.NewMenuItem(sf)
	sf.N4.SetCaption("关于")
	sf.N4.SetOnClick(sf.N4Click)

	sf.Add(sf.N4)

	return sf
}

func (sf *TMenuControl) MenuControlStartClick(sender vcl.IObject) {
	FrmMain.startService()
}

func (sf *TMenuControl) MenuControlStopClick(sender vcl.IObject) {
	if vcl.MessageDlg("是否确认停止服务?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.MrYes {
		FrmMain.stopService()
	}
}

func (sf *TMenuControl) MenuControlReconnectClick(sender vcl.IObject) {
	ReConnectServerTick = 0
}

func (sf *TMenuControl) MenuControlClearLogClick(sender vcl.IObject) {
	if vcl.MessageDlg("是否确认清除显示的日志信息?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.MrYes {
		FrmMain.MemoLog.Clear()
	}
}

func (sf *TMenuControl) MenuControlExitClick(sender vcl.IObject) {
	FrmMain.Close()
}

func (sf *TMenuView) MenuViewLogMsgClick(sender vcl.IObject) {
	sf.MenuViewLogMsg.SetChecked(!sf.MenuViewLogMsg.Checked())
	FrmMain.showLogMsg(sf.MenuViewLogMsg.Checked())
}

func (sf *TMenuOption) MenuOptionGeneralClick(sender vcl.IObject) {
	FrmGeneralConfig.ShowModal()
}

func (sf *TMenuOption) MenuOptionIpFilterClick(sender vcl.IObject) {
	FrmIPAddrFilter.ShowModal()
}

func (sf *TMenuItem3) N4Click(sender vcl.IObject) {
	MainLog.AddMsg("引擎版本: 1.5.0 (20020522)")
	MainLog.AddMsg("更新日期: 2023/09/14")
	MainLog.AddMsg("程序制作: CHUNQIAN SHEN")
}
