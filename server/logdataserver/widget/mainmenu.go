// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package widget

import (
	"github.com/ying32/govcl/vcl"
)

type TMainMenu struct {
	*vcl.TMainMenu

	menuControl *TMenuControl
	menuHelp    *TMenuHelp
}

func NewMainMenu(sender vcl.IComponent) *TMainMenu {
	sf := &TMainMenu{
		TMainMenu: vcl.NewMainMenu(sender),
	}

	sf.menuControl = NewMenuControl(sf)
	sf.menuControl.SetCaption("控制")

	sf.menuHelp = NewMenuHelp(sf)
	sf.menuHelp.SetCaption("帮助")

	sf.Items().Add(sf.menuControl)
	sf.Items().Add(sf.menuHelp)

	return sf
}
