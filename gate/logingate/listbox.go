// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import "github.com/ying32/govcl/vcl"

type TListBoxActiveList struct {
	*vcl.TListBox

	ActiveListPopupMenu *TActiveListPopupMenu
}

type TListBoxBlockList struct {
	*vcl.TListBox

	BlockListPopupMenu *TBlockListPopupMenu
}

type TListBoxTempList struct {
	*vcl.TListBox

	TempBlockListPopupMenu *TTempBlockListPopupMenu
}

func NewListBoxActiveList(sender vcl.IComponent) *TListBoxActiveList {
	sf := &TListBoxActiveList{
		TListBox: vcl.NewListBox(sender),
	}

	sf.ActiveListPopupMenu = NewActiveListPopupMenu(sf)

	sf.SetPopupMenu(sf.ActiveListPopupMenu)

	return sf
}

func NewListBoxBlockList(sender vcl.IComponent) *TListBoxBlockList {
	sf := &TListBoxBlockList{
		TListBox: vcl.NewListBox(sender),
	}

	sf.BlockListPopupMenu = NewBlockListPopupMenu(sf)

	sf.SetPopupMenu(sf.BlockListPopupMenu)
	return sf
}

func NewListBoxTempList(sender vcl.IComponent) *TListBoxTempList {
	sf := &TListBoxTempList{
		TListBox: vcl.NewListBox(sender),
	}

	sf.TempBlockListPopupMenu = NewTempBlockListPopupMenu(sf)

	sf.SetPopupMenu(sf.TempBlockListPopupMenu)

	return sf
}
