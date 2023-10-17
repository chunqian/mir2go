// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package widget

import "github.com/ying32/govcl/vcl"

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
	// sf.menuControlClearLog.SetOnClick(sf.menuControlClearLogClick)

	sf.Add(sf.menuControlClearLog)

	return sf
}

func NewMenuHelp(sender vcl.IComponent) *TMenuHelp {
	sf := &TMenuHelp{
		TMenuItem: vcl.NewMenuItem(sender),
	}

	sf.menuAbout = vcl.NewMenuItem(sf)
	sf.menuAbout.SetCaption("关于")
	// sf.menuAbout.SetOnClick(sf.menuAboutClick)

	sf.Add(sf.menuAbout)

	return sf
}
