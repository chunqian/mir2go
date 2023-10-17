// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package widget

import "github.com/ying32/govcl/vcl"

type TListBoxActiveList struct {
	*vcl.TListBox

	activeListPopupMenu *TActiveListPopupMenu
}

type TListBoxBlockList struct {
	*vcl.TListBox

	blockListPopupMenu *TBlockListPopupMenu
}

type TListBoxTempList struct {
	*vcl.TListBox

	tempBlockListPopupMenu *TTempBlockListPopupMenu
}

func NewListBoxActiveList(sender vcl.IComponent) *TListBoxActiveList {
	sf := &TListBoxActiveList{
		TListBox: vcl.NewListBox(sender),
	}

	sf.activeListPopupMenu = NewActiveListPopupMenu(sf)

	sf.SetPopupMenu(sf.activeListPopupMenu)

	return sf
}

func NewListBoxBlockList(sender vcl.IComponent) *TListBoxBlockList {
	sf := &TListBoxBlockList{
		TListBox: vcl.NewListBox(sender),
	}

	sf.blockListPopupMenu = NewBlockListPopupMenu(sf)

	sf.SetPopupMenu(sf.blockListPopupMenu)
	return sf
}

func NewListBoxTempList(sender vcl.IComponent) *TListBoxTempList {
	sf := &TListBoxTempList{
		TListBox: vcl.NewListBox(sender),
	}

	sf.tempBlockListPopupMenu = NewTempBlockListPopupMenu(sf)

	sf.SetPopupMenu(sf.tempBlockListPopupMenu)

	return sf
}
