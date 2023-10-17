// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package widget

import (
	"github.com/ying32/govcl/vcl"
)

type TMainMenu struct {
	*vcl.TMainMenu

	menuControl *TMenuControl
	menuView    *TMenuView
	menuOption  *TMenuOption
	n3          *TMenuItem3
}

func NewMainMenu(sender vcl.IComponent) *TMainMenu {
	sf := &TMainMenu{
		TMainMenu: vcl.NewMainMenu(sender),
	}

	sf.menuControl = NewMenuControl(sf)
	sf.menuControl.SetCaption("控制")

	sf.menuView = NewMenuView(sf)
	sf.menuView.SetCaption("查看")

	sf.menuOption = NewMenuOption(sf)
	sf.menuOption.SetCaption("选项")

	sf.n3 = NewMenuItem3(sf)
	sf.n3.SetCaption("帮助")

	sf.Items().Add(sf.menuControl)
	sf.Items().Add(sf.menuView)
	sf.Items().Add(sf.menuOption)
	sf.Items().Add(sf.n3)

	return sf
}
