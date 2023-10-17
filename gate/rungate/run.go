// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	"github.com/chunqian/mir2go/gate/rungate/widget"
)

type TFrmMain struct {
	*vcl.TForm

	mainMenu     *widget.TMainMenu
	pageControl1 *widget.TPageControl
	statusBar    *vcl.TStatusBar
}

var (
	frmMain *TFrmMain
)

func (sf *TFrmMain) SetComponents() {
	sf.mainMenu = widget.NewMainMenu(sf)
}

func (sf *TFrmMain) OnFormCreate(sender vcl.IObject) {
	// 布局
	sf.SetCaption("游戏网关")
	sf.EnabledMaximize(false)
	sf.SetBorderStyle(types.BsSingle)
	sf.SetBounds(433, 275, 308, 197)
	sf.SetComponents()
}

func (sf *TFrmMain) OnFormDestroy(sender vcl.IObject) {

}

func (sf *TFrmMain) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	*canClose = vcl.MessageDlg("是否确认退出服务器?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.MrYes
}
