// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
)

type TFrmIPaddrFilter struct {
	*vcl.TForm

	ButtonOK       *vcl.TButton
	GroupBox1      *TGroupBox1
	GroupBox2      *TGroupBox2
	GroupBoxActive *TGroupBoxActive
}

type TActiveListPopupMenu struct {
	*vcl.TPopupMenu

	PopMenuRefList     *vcl.TMenuItem
	PopMenuSort        *vcl.TMenuItem
	PopMenuAddTempList *vcl.TMenuItem
	PopMenuBlockList   *vcl.TMenuItem
	PopMenuKick        *vcl.TMenuItem
}

type TBlockListPopupMenu struct {
	*vcl.TPopupMenu

	PopMenuRefList     *vcl.TMenuItem
	PopMenuSort        *vcl.TMenuItem
	PopMenuAdd         *vcl.TMenuItem
	PopMenuAddTempList *vcl.TMenuItem
	PopMenuDelete      *vcl.TMenuItem
}

type TTempBlockListPopupMenu struct {
	*vcl.TPopupMenu

	PopMenuRefList   *vcl.TMenuItem
	PopMenuSort      *vcl.TMenuItem
	PopMenuAdd       *vcl.TMenuItem
	PopMenuBlockList *vcl.TMenuItem
	PopMenuDelete    *vcl.TMenuItem
}

type TGroupBox1 struct {
	*vcl.TGroupBox

	LabelTempList    *vcl.TLabel
	Label1           *vcl.TLabel
	ListBoxTempList  *TListBoxTempList
	ListBoxBlockList *TListBoxBlockList
}

type TGroupBox2 struct {
	*vcl.TGroupBox

	EditMaxConnect *vcl.TSpinEdit
	GroupBox3      *TGroupBox3
	Label2         *vcl.TLabel
	Label3         *vcl.TLabel
	Label7         *vcl.TLabel
}

type TGroupBox3 struct {
	*vcl.TGroupBox

	RadioAddBlockList *vcl.TRadioButton
	RadioAddTempList  *vcl.TRadioButton
	RadioDisConnect   *vcl.TRadioButton
}

type TGroupBoxActive struct {
	*vcl.TGroupBox

	Label4            *vcl.TLabel
	ListBoxActiveList *TListBoxActiveList
}

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

var (
	FrmIPaddrFilter *TFrmIPaddrFilter
)

func (sf *TFrmIPaddrFilter) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.Layout()
}

func (sf *TFrmIPaddrFilter) Layout() {

	sf.SetLeft(420)
	sf.SetTop(296)
	sf.SetClientWidth(679)
	sf.SetClientHeight(367)

	// sf.BlockListPopupMenu = &TBlockListPopupMenu{
	// 	TPopupMenu: vcl.NewPopupMenu(sf),
	// }
	// sf.BlockListPopupMenu.Layout(sf)

	sf.ButtonOK = vcl.NewButton(sf)
	sf.ButtonOK.SetCaption("确定(&O)")
	sf.ButtonOK.SetDefault(true)
	sf.ButtonOK.SetLeft(568)
	sf.ButtonOK.SetTop(295)
	sf.ButtonOK.SetWidth(86)
	sf.ButtonOK.SetHeight(27)
	sf.ButtonOK.SetParent(sf)

	sf.GroupBoxActive = &TGroupBoxActive{
		TGroupBox: vcl.NewGroupBox(sf),
	}

	sf.GroupBoxActive.SetCaption("当前连接")
	sf.GroupBoxActive.Font().SetSize(9)
	sf.GroupBoxActive.SetLeft(9)
	sf.GroupBoxActive.SetTop(9)
	sf.GroupBoxActive.SetWidth(148)
	sf.GroupBoxActive.SetHeight(313)
	sf.GroupBoxActive.Layout(sf)
	sf.GroupBoxActive.SetParent(sf)

	sf.GroupBox1 = &TGroupBox1{
		TGroupBox: vcl.NewGroupBox(sf),
	}
	sf.GroupBox1.SetCaption("过滤列表")
	sf.GroupBox1.Font().SetSize(9)
	sf.GroupBox1.SetLeft(162)
	sf.GroupBox1.SetTop(9)
	sf.GroupBox1.SetWidth(294)
	sf.GroupBox1.SetHeight(313)
	sf.GroupBox1.Layout(sf)
	sf.GroupBox1.SetParent(sf)

	sf.GroupBox2 = &TGroupBox2{
		TGroupBox: vcl.NewGroupBox(sf),
	}
	sf.GroupBox2.SetCaption("攻击保护")
	sf.GroupBox2.Font().SetSize(9)
	sf.GroupBox2.SetLeft(464)
	sf.GroupBox2.SetTop(9)
	sf.GroupBox2.SetWidth(201)
	sf.GroupBox2.SetHeight(274)
	sf.GroupBox2.Layout(sf)
	sf.GroupBox2.SetParent(sf)
}

func (sf *TActiveListPopupMenu) Layout(sender *TFrmIPaddrFilter) {

	sf.PopMenuRefList = vcl.NewMenuItem(sf)
	sf.PopMenuRefList.SetCaption("刷新(&R)")

	sf.PopMenuSort = vcl.NewMenuItem(sf)
	sf.PopMenuSort.SetCaption("排序(&S)")

	sf.PopMenuAddTempList = vcl.NewMenuItem(sf)
	sf.PopMenuAddTempList.SetCaption("加入动态过滤列表(&A)")

	sf.PopMenuBlockList = vcl.NewMenuItem(sf)
	sf.PopMenuBlockList.SetCaption("加入永久过滤列表(&D)")

	sf.PopMenuKick = vcl.NewMenuItem(sf)
	sf.PopMenuKick.SetCaption("踢除下线(&K)")

	sf.Items().Add(sf.PopMenuRefList)
	sf.Items().Add(sf.PopMenuSort)
	sf.Items().Add(sf.PopMenuAddTempList)
	sf.Items().Add(sf.PopMenuBlockList)
	sf.Items().Add(sf.PopMenuKick)
}

func (sf *TBlockListPopupMenu) Layout(sender *TFrmIPaddrFilter) {
	sf.PopMenuRefList = vcl.NewMenuItem(sf)
	sf.PopMenuRefList.SetCaption("刷新(&R)")

	sf.PopMenuSort = vcl.NewMenuItem(sf)
	sf.PopMenuSort.SetCaption("排序(&S)")

	sf.PopMenuAdd = vcl.NewMenuItem(sf)
	sf.PopMenuAdd.SetCaption("增加(&A)")

	sf.PopMenuAddTempList = vcl.NewMenuItem(sf)
	sf.PopMenuAddTempList.SetCaption("加入动态过滤列表(&A)")

	sf.PopMenuDelete = vcl.NewMenuItem(sf)
	sf.PopMenuDelete.SetCaption("删除(&D)")

	sf.Items().Add(sf.PopMenuRefList)
	sf.Items().Add(sf.PopMenuSort)
	sf.Items().Add(sf.PopMenuAdd)
	sf.Items().Add(sf.PopMenuAddTempList)
	sf.Items().Add(sf.PopMenuDelete)
}

func (sf *TTempBlockListPopupMenu) Layout(sender *TFrmIPaddrFilter) {
	sf.PopMenuRefList = vcl.NewMenuItem(sf)
	sf.PopMenuRefList.SetCaption("刷新(&R)")

	sf.PopMenuSort = vcl.NewMenuItem(sf)
	sf.PopMenuSort.SetCaption("排序(&S)")

	sf.PopMenuAdd = vcl.NewMenuItem(sf)
	sf.PopMenuAdd.SetCaption("增加(&A)")

	sf.PopMenuBlockList = vcl.NewMenuItem(sf)
	sf.PopMenuBlockList.SetCaption("加入永久过滤列表(&D)")

	sf.PopMenuDelete = vcl.NewMenuItem(sf)
	sf.PopMenuDelete.SetCaption("删除(&D)")

	sf.Items().Add(sf.PopMenuRefList)
	sf.Items().Add(sf.PopMenuSort)
	sf.Items().Add(sf.PopMenuAdd)
	sf.Items().Add(sf.PopMenuBlockList)
	sf.Items().Add(sf.PopMenuDelete)
}

func (sf *TGroupBoxActive) Layout(sender *TFrmIPaddrFilter) {

	sf.Label4 = vcl.NewLabel(sf)
	sf.Label4.SetCaption("连接列表:")
	sf.Label4.SetLeft(0)
	sf.Label4.SetTop(9)
	sf.Label4.SetWidth(59)
	sf.Label4.SetHeight(13)
	sf.Label4.SetParent(sf)

	sf.ListBoxActiveList = &TListBoxActiveList{
		TListBox: vcl.NewListBox(sf),
	}
	sf.ListBoxActiveList.SetHint("当前连接的IP地址列表")
	sf.ListBoxActiveList.SetLeft(0)
	sf.ListBoxActiveList.SetTop(31)
	sf.ListBoxActiveList.SetWidth(138)
	sf.ListBoxActiveList.SetHeight(261)
	sf.ListBoxActiveList.SetItemHeight(13)
	sf.ListBoxActiveList.SetParentShowHint(false)
	sf.ListBoxActiveList.SetShowHint(true)
	sf.ListBoxActiveList.SetSorted(true)
	sf.ListBoxActiveList.Layout(sender)
	sf.ListBoxActiveList.SetParent(sf)
}

func (sf *TListBoxActiveList) Layout(sender *TFrmIPaddrFilter) {

	sf.ActiveListPopupMenu = &TActiveListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(sf),
	}
	sf.ActiveListPopupMenu.Layout(sender)

	sf.SetPopupMenu(sf.ActiveListPopupMenu)
}

func (sf *TGroupBox1) Layout(sender *TFrmIPaddrFilter) {

	sf.LabelTempList = vcl.NewLabel(sf)
	sf.LabelTempList.SetCaption("动态过滤:")
	sf.LabelTempList.SetLeft(9)
	sf.LabelTempList.SetTop(9)
	sf.LabelTempList.SetWidth(59)
	sf.LabelTempList.SetHeight(13)
	sf.LabelTempList.SetParent(sf)

	sf.Label1 = vcl.NewLabel(sf)
	sf.Label1.SetCaption("永久过滤:")
	sf.Label1.SetLeft(147)
	sf.Label1.SetTop(9)
	sf.Label1.SetWidth(59)
	sf.Label1.SetHeight(13)
	sf.Label1.SetParent(sf)

	sf.ListBoxTempList = &TListBoxTempList{
		TListBox: vcl.NewListBox(sf),
	}
	sf.ListBoxTempList.SetHint("动态过滤列表,在此列表中的IP将无法建立连接,但在程序重新启动时此列表的信息将被清空")
	sf.ListBoxTempList.SetLeft(0)
	sf.ListBoxTempList.SetTop(31)
	sf.ListBoxTempList.SetWidth(138)
	sf.ListBoxTempList.SetHeight(261)
	sf.ListBoxTempList.SetItemHeight(13)
	sf.ListBoxTempList.SetParentShowHint(false)
	sf.ListBoxTempList.SetShowHint(true)
	sf.ListBoxTempList.SetSorted(true)
	sf.ListBoxTempList.Layout(sender)
	sf.ListBoxTempList.SetParent(sf)

	sf.ListBoxBlockList = &TListBoxBlockList{
		TListBox: vcl.NewListBox(sf),
	}
	sf.ListBoxBlockList.SetHint("永久过滤列表,在此列表中的IP将无法建立连接,此列表将保存于配置文件中,在程序重新启动时会重新加载此列表")
	sf.ListBoxBlockList.SetLeft(147)
	sf.ListBoxBlockList.SetTop(31)
	sf.ListBoxBlockList.SetWidth(138)
	sf.ListBoxBlockList.SetHeight(261)
	sf.ListBoxBlockList.SetItemHeight(13)
	sf.ListBoxBlockList.SetParentShowHint(false)
	sf.ListBoxBlockList.SetShowHint(true)
	sf.ListBoxBlockList.SetSorted(true)
	sf.ListBoxBlockList.Layout(sender)
	sf.ListBoxBlockList.SetParent(sf)
}

func (sf *TListBoxBlockList) Layout(sender *TFrmIPaddrFilter) {

	sf.BlockListPopupMenu = &TBlockListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(sf),
	}
	sf.BlockListPopupMenu.Layout(sender)

	sf.SetPopupMenu(sf.BlockListPopupMenu)
}

func (sf *TListBoxTempList) Layout(sender *TFrmIPaddrFilter) {

	sf.TempBlockListPopupMenu = &TTempBlockListPopupMenu{
		TPopupMenu: vcl.NewPopupMenu(sf),
	}
	sf.TempBlockListPopupMenu.Layout(sender)

	sf.SetPopupMenu(sf.TempBlockListPopupMenu)
}

func (sf *TGroupBox2) Layout(sender *TFrmIPaddrFilter) {
}
