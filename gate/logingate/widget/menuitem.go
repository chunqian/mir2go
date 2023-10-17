// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package widget

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	. "github.com/chunqian/mir2go/common"
)

type TMenuControl struct {
	*vcl.TMenuItem

	menuControlStart     *vcl.TMenuItem
	menuControlStop      *vcl.TMenuItem
	menuControlReconnect *vcl.TMenuItem
	n1                   *vcl.TMenuItem
	menuControlClearLog  *vcl.TMenuItem
	n2                   *vcl.TMenuItem
	menuControlExit      *vcl.TMenuItem
}

type TMenuView struct {
	*vcl.TMenuItem

	menuViewLogMsg *vcl.TMenuItem
}

type TMenuOption struct {
	*vcl.TMenuItem

	menuOptionGeneral  *vcl.TMenuItem
	menuOptionIpFilter *vcl.TMenuItem
}

type TMenuItem3 struct {
	*vcl.TMenuItem

	n4 *vcl.TMenuItem
}

func NewMenuControl(sender vcl.IComponent) *TMenuControl {
	sf := &TMenuControl{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.menuControlStart = vcl.NewMenuItem(sf)
	sf.menuControlStart.SetCaption("启动服务")
	sf.menuControlStart.SetShortCutFromString("Ctrl+S")
	sf.menuControlStart.SetOnClick(sf.menuControlStartClick)

	sf.menuControlStop = vcl.NewMenuItem(sf)
	sf.menuControlStop.SetCaption("停止服务")
	sf.menuControlStop.SetShortCutFromString("Ctrl+T")
	sf.menuControlStop.SetOnClick(sf.menuControlStopClick)

	sf.menuControlReconnect = vcl.NewMenuItem(sf)
	sf.menuControlReconnect.SetCaption("刷新连接")
	sf.menuControlReconnect.SetShortCutFromString("Ctrl+R")
	sf.menuControlReconnect.SetOnClick(sf.menuControlReconnectClick)

	sf.menuControlClearLog = vcl.NewMenuItem(sf)
	sf.menuControlClearLog.SetCaption("清除日志")
	sf.menuControlClearLog.SetShortCutFromString("Ctrl+C")
	sf.menuControlClearLog.SetOnClick(sf.menuControlClearLogClick)

	sf.menuControlExit = vcl.NewMenuItem(sf)
	sf.menuControlExit.SetCaption("退出")
	sf.menuControlExit.SetShortCutFromString("Ctrl+X")
	sf.menuControlExit.SetOnClick(sf.menuControlExitClick)

	sf.Add(sf.menuControlStart)
	sf.Add(sf.menuControlStop)
	sf.Add(sf.menuControlReconnect)
	sf.Add(sf.n1)
	sf.Add(sf.menuControlClearLog)
	sf.Add(sf.n2)
	sf.Add(sf.menuControlExit)

	// 注册观察者
	GetSubject("widget.TMainMenu.TMenuControl").Register(sf)

	return sf
}

func NewMenuView(sender vcl.IComponent) *TMenuView {
	sf := &TMenuView{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.menuViewLogMsg = vcl.NewMenuItem(sf)
	sf.menuViewLogMsg.SetCaption("查看日志")
	sf.menuViewLogMsg.SetOnClick(sf.menuViewLogMsgClick)

	sf.Add(sf.menuViewLogMsg)

	// 注册观察者
	GetSubject("widget.TMainMenu.TMenuView").Register(sf)

	return sf
}

func NewMenuOption(sender vcl.IComponent) *TMenuOption {
	sf := &TMenuOption{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.menuOptionGeneral = vcl.NewMenuItem(sf)
	sf.menuOptionGeneral.SetCaption("基本设置")
	sf.menuOptionGeneral.SetOnClick(sf.menuOptionGeneralClick)

	sf.menuOptionIpFilter = vcl.NewMenuItem(sf)
	sf.menuOptionIpFilter.SetCaption("安全过滤")
	sf.menuOptionIpFilter.SetOnClick(sf.menuOptionIpFilterClick)

	sf.Add(sf.menuOptionGeneral)
	sf.Add(sf.menuOptionIpFilter)

	return sf
}

func NewMenuItem3(sender vcl.IComponent) *TMenuItem3 {
	sf := &TMenuItem3{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.n4 = vcl.NewMenuItem(sf)
	sf.n4.SetCaption("关于")
	sf.n4.SetOnClick(sf.n4Click)

	sf.Add(sf.n4)

	return sf
}

func (sf *TMenuControl) menuControlStartClick(sender vcl.IObject) {
	GetSubject("TFrmMain").Notify("menuControlStartClick", nil)
}

func (sf *TMenuControl) menuControlStopClick(sender vcl.IObject) {
	if vcl.MessageDlg("是否确认停止服务?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.MrYes {
		GetSubject("TFrmMain").Notify("menuControlStopClick", nil)
	}
}

func (sf *TMenuControl) menuControlReconnectClick(sender vcl.IObject) {
	GetSubject("TFrmMain").Notify("menuControlReconnectClick", nil)
}

func (sf *TMenuControl) menuControlClearLogClick(sender vcl.IObject) {
	if vcl.MessageDlg("是否确认清除显示的日志信息?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.MrYes {
		GetSubject("TFrmMain").Notify("menuControlClearLogClick", nil)
	}
}

func (sf *TMenuControl) menuControlExitClick(sender vcl.IObject) {
	// 通知所有观察者
	GetSubject("TFrmMain").Notify("menuControlExitClick", nil)
}

func (sf *TMenuControl) ObserverNotifyReceived(tag string, data interface{}) {
	switch tag {
	case "SetMenuControlStart":
		sf.SetEnabled(data.(bool))
	case "SetMenuControlStop":
		sf.SetEnabled(data.(bool))
	}
}

func (sf *TMenuView) menuViewLogMsgClick(sender vcl.IObject) {
	sf.menuViewLogMsg.SetChecked(!sf.menuViewLogMsg.Checked())
	GetSubject("TFrmMain").Notify("menuViewLogMsgClick", sf.menuViewLogMsg.Checked())
}

func (sf *TMenuView) ObserverNotifyReceived(tag string, data interface{}) {
	switch tag {
	case "MenuViewLogMsgClick":
		sf.menuViewLogMsgClick(sf)
	}
}

func (sf *TMenuOption) menuOptionGeneralClick(sender vcl.IObject) {
	GetSubject("TFrmGeneralConfig").Notify("menuOptionGeneralClick", nil)
}

func (sf *TMenuOption) menuOptionIpFilterClick(sender vcl.IObject) {
	GetSubject("TFrmIPAddrFilter").Notify("menuOptionIpFilterClick", nil)
}

func (sf *TMenuItem3) n4Click(sender vcl.IObject) {
	MainLog.AddMsg("引擎版本: 1.5.0 (20020522)")
	MainLog.AddMsg("更新日期: 2023/09/14")
	MainLog.AddMsg("程序制作: CHUNQIAN SHEN")
}
