// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import "github.com/ying32/govcl/vcl"

// ******************** TMainMenu ********************
type TMainMenu struct {
	*vcl.TMainMenu

	MenuControl *TMenuControl
	MenuView    *TMenuView
	MenuOption  *TMenuOption
	N3          *TMenuItem3
}

func NewMainMenu(sender vcl.IComponent) *TMainMenu {
	sf := &TMainMenu{
		TMainMenu: vcl.NewMainMenu(sender),
	}

	sf.MenuControl = NewMenuControl(sf)
	sf.MenuControl.SetCaption("控制")

	sf.MenuView = NewMenuView(sf)
	sf.MenuView.SetCaption("查看")

	sf.MenuOption = NewMenuOption(sf)
	sf.MenuOption.SetCaption("选项")

	sf.N3 = NewMenuItem3(sf)
	sf.N3.SetCaption("帮助")

	sf.Items().Add(sf.MenuControl)
	sf.Items().Add(sf.MenuView)
	sf.Items().Add(sf.MenuOption)
	sf.Items().Add(sf.N3)

	return sf
}
