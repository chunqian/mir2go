// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package widget

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	. "github.com/chunqian/mir2go/common"
)

type TMenuControl struct {
	*vcl.TMenuItem

	menuControlClearLog *vcl.TMenuItem
}

type TMenuHelp struct {
	*vcl.TMenuItem

	menuAbout *vcl.TMenuItem
}

func NewMenuControl(sender vcl.IComponent) *TMenuControl {
	sf := &TMenuControl{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.menuControlClearLog = vcl.NewMenuItem(sf)
	sf.menuControlClearLog.SetCaption("清除日志")
	sf.menuControlClearLog.SetShortCutFromString("Ctrl+C")
	sf.menuControlClearLog.SetOnClick(sf.menuControlClearLogClick)

	sf.Add(sf.menuControlClearLog)

	return sf
}

func NewMenuHelp(sender vcl.IComponent) *TMenuHelp {
	sf := &TMenuHelp{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.menuAbout = vcl.NewMenuItem(sf)
	sf.menuAbout.SetCaption("关于")
	sf.menuAbout.SetOnClick(sf.menuAboutClick)

	sf.Add(sf.menuAbout)

	return sf
}

func (sf *TMenuControl) menuControlClearLogClick(sender vcl.IObject) {
	if vcl.MessageDlg("是否确认清除显示的日志信息?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.MrYes {
		// 通知
		ObserverGetTopic("TFrmLogData").Notify("menuControlClearLogClick", nil)
	}
}

func (sf *TMenuHelp) menuAboutClick(sender vcl.IObject) {
	MainLog.AddMsg("引擎版本: 1.5.0 (20020522)")
	MainLog.AddMsg("更新日期: 2023/09/14")
	MainLog.AddMsg("程序制作: CHUNQIAN SHEN")
	// 通知
	ObserverGetTopic("TFrmLogData").Notify("showMainLogMsg", nil)
}
