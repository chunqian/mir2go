// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import "github.com/ying32/govcl/vcl"

type TFrmIPaddrFilter struct {
	*vcl.TForm

	ActiveListPopupMenu *TActiveListPopupMenu
	BlockListPopupMenu *TBlockListPopupMenu
	ButtonOK *vcl.TButton
	GroupBox1 *TGroupBox1
	GroupBox2 *TGroupBox2
	GroupBoxActive *TGroupBoxActive
	TempBlockListPopupMenu *TTempBlockListPopupMenu
}

type TActiveListPopupMenu struct {
	*vcl.TPopupMenu

	APOPMENU_REFLIST *vcl.TMenuItem
	APOPMENU_SORT *vcl.TMenuItem
	APOPMENU_ADDTEMPLIST *vcl.TMenuItem
	APOPMENU_BLOCKLIST *vcl.TMenuItem
	APOPMENU_KICK *vcl.TMenuItem
}

type TBlockListPopupMenu struct {
	*vcl.TPopupMenu

	BPOPMENU_REFLIST *vcl.TMenuItem
	BPOPMENU_SORT *vcl.TMenuItem
	BPOPMENU_ADD *vcl.TMenuItem
	BPOPMENU_ADDTEMPLIST *vcl.TMenuItem
	BPOPMENU_DELETE *vcl.TMenuItem
}

type TTempBlockListPopupMenu struct {
	*vcl.TPopupMenu

	TPOPMENU_REFLIST *vcl.TMenuItem
	TPOPMENU_SORT *vcl.TMenuItem
	TPOPMENU_ADD *vcl.TMenuItem
	TPOPMENU_BLOCKLIST *vcl.TMenuItem
	TPOPMENU_DELETE *vcl.TMenuItem
}

type TGroupBox1 struct {
	*vcl.TGroupBox

	Label1 *vcl.TLabel
	LabelTempList *vcl.TLabel
	ListBoxBlockList *vcl.TListBox
	ListBoxTempList *vcl.TListBox
}

type TGroupBox2 struct {
	*vcl.TGroupBox

	EditMaxConnect * vcl.TSpinEdit
	GroupBox3 *TGroupBox3
	Label2 *vcl.TLabel
	Label3 *vcl.TLabel
	Label7 *vcl.TLabel
}

type TGroupBox3 struct {
	*vcl.TGroupBox

	RadioAddBlockList *vcl.TRadioButton
	RadioAddTempList *vcl.TRadioButton
	RadioDisConnect *vcl.TRadioButton
}

type TGroupBoxActive struct {
	*vcl.TGroupBox

	Label4 *vcl.TLabel
	ListBoxActiveList *vcl.TListBox
}

var (
	FrmIPaddrFilter *TFrmIPaddrFilter
)

func (frm *TFrmIPaddrFilter) OnFormCreate(sender vcl.IObject) {
	frm.layout()
}

func (frm *TFrmIPaddrFilter) layout() {
	frm.ActiveListPopupMenu = &TActiveListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(frm),
	}
	frm.ActiveListPopupMenu.layout()

	frm.BlockListPopupMenu = &TBlockListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(frm),
	}
	frm.BlockListPopupMenu.layout()

	frm.ButtonOK = vcl.NewButton(frm)
	frm.ButtonOK.SetCaption("确定(&O)")
	frm.ButtonOK.SetDefault(true)
	frm.ButtonOK.SetHeight(27)
	frm.ButtonOK.SetLeft(568)
	frm.ButtonOK.SetTop(295)
	frm.ButtonOK.SetWidth(86)
	frm.ButtonOK.SetParent(frm)
}

func (act *TActiveListPopupMenu) layout() {
	act.APOPMENU_REFLIST = vcl.NewMenuItem(act)
	act.APOPMENU_REFLIST.SetCaption("刷新(&R)")

	act.APOPMENU_SORT = vcl.NewMenuItem(act)
	act.APOPMENU_SORT.SetCaption("排序(&S)")

	act.APOPMENU_ADDTEMPLIST = vcl.NewMenuItem(act)
	act.APOPMENU_ADDTEMPLIST.SetCaption("加入动态过滤列表(&A)")

	act.APOPMENU_BLOCKLIST = vcl.NewMenuItem(act)
	act.APOPMENU_BLOCKLIST.SetCaption("加入永久过滤列表(&D)")

	act.APOPMENU_KICK = vcl.NewMenuItem(act)
	act.APOPMENU_KICK.SetCaption("踢除下线(&K)")
}

func (blk *TBlockListPopupMenu) layout() {
	blk.BPOPMENU_REFLIST = vcl.NewMenuItem(blk)
	blk.BPOPMENU_REFLIST.SetCaption("刷新(&R)")

	blk.BPOPMENU_SORT = vcl.NewMenuItem(blk)
	blk.BPOPMENU_SORT.SetCaption("排序(&S)")

	blk.BPOPMENU_ADD = vcl.NewMenuItem(blk)
	blk.BPOPMENU_ADD.SetCaption("增加(&A)")

	blk.BPOPMENU_ADDTEMPLIST = vcl.NewMenuItem(blk)
	blk.BPOPMENU_ADDTEMPLIST.SetCaption("加入动态过滤列表(&A)")

	blk.BPOPMENU_DELETE = vcl.NewMenuItem(blk)
	blk.BPOPMENU_DELETE.SetCaption("删除(&D)")
}
