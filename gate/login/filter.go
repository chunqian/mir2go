// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
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

// ******************** Layout ********************
func (sf *TFrmIPaddrFilter) Layout() {

	sf.ButtonOK = vcl.NewButton(sf)
	sf.ButtonOK.SetCaption("确定(&O)")
	sf.ButtonOK.SetDefault(true)
	sf.ButtonOK.SetBounds(568, 295, 86, 27)

	sf.GroupBoxActive = &TGroupBoxActive{
		TGroupBox: vcl.NewGroupBox(sf),
	}
	sf.GroupBoxActive.SetCaption("当前连接")
	sf.GroupBoxActive.Font().SetSize(9)
	sf.GroupBoxActive.SetBounds(9, 9, 148, 313)
	sf.GroupBoxActive.Layout(sf)

	sf.GroupBox1 = &TGroupBox1{
		TGroupBox: vcl.NewGroupBox(sf),
	}
	sf.GroupBox1.SetCaption("过滤列表")
	sf.GroupBox1.Font().SetSize(9)
	sf.GroupBox1.SetBounds(162, 9, 294, 313)
	sf.GroupBox1.Layout(sf)

	sf.GroupBox2 = &TGroupBox2{
		TGroupBox: vcl.NewGroupBox(sf),
	}
	sf.GroupBox2.SetCaption("攻击保护")
	sf.GroupBox2.Font().SetSize(9)
	sf.GroupBox2.SetBounds(464, 9, 201, 274)
	sf.GroupBox2.Layout(sf)

	sf.ButtonOK.SetParent(sf)
	sf.GroupBoxActive.SetParent(sf)
	sf.GroupBox1.SetParent(sf)
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
	sf.Label4.SetBounds(0, 9, 59, 13)

	sf.ListBoxActiveList = &TListBoxActiveList{
		TListBox: vcl.NewListBox(sf),
	}
	sf.ListBoxActiveList.SetHint("当前连接的IP地址列表")
	sf.ListBoxActiveList.SetItemHeight(13)
	sf.ListBoxActiveList.SetBounds(0, 31, 138, 261)
	sf.ListBoxActiveList.SetParentShowHint(false)
	sf.ListBoxActiveList.SetShowHint(true)
	sf.ListBoxActiveList.SetSorted(true)
	sf.ListBoxActiveList.Layout(sender)

	sf.Label4.SetParent(sf)
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
	sf.LabelTempList.SetBounds(9, 9, 59, 13)

	sf.Label1 = vcl.NewLabel(sf)
	sf.Label1.SetCaption("永久过滤:")
	sf.Label1.SetBounds(147, 9, 59, 13)

	sf.ListBoxTempList = &TListBoxTempList{
		TListBox: vcl.NewListBox(sf),
	}
	sf.ListBoxTempList.SetHint("动态过滤列表, 在此列表中的IP将无法建立连接, 但在程序重新启动时此列表的信息将被清空")
	sf.ListBoxTempList.SetItemHeight(13)
	sf.ListBoxTempList.SetBounds(0, 31, 138, 261)
	sf.ListBoxTempList.SetParentShowHint(false)
	sf.ListBoxTempList.SetShowHint(true)
	sf.ListBoxTempList.SetSorted(true)
	sf.ListBoxTempList.Layout(sender)

	sf.ListBoxBlockList = &TListBoxBlockList{
		TListBox: vcl.NewListBox(sf),
	}
	sf.ListBoxBlockList.SetHint("永久过滤列表, 在此列表中的IP将无法建立连接, 此列表将保存于配置文件中, 在程序重新启动时会重新加载此列表")
	sf.ListBoxBlockList.SetItemHeight(13)
	sf.ListBoxBlockList.SetBounds(147, 31, 138, 261)
	sf.ListBoxBlockList.SetParentShowHint(false)
	sf.ListBoxBlockList.SetShowHint(true)
	sf.ListBoxBlockList.SetSorted(true)
	sf.ListBoxBlockList.Layout(sender)

	sf.LabelTempList.SetParent(sf)
	sf.Label1.SetParent(sf)
	sf.ListBoxTempList.SetParent(sf)
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

	sf.Label2 = vcl.NewLabel(sf)
	sf.Label2.SetCaption("连接限制:")
	sf.Label2.SetBounds(9, 9, 59, 13)

	sf.EditMaxConnect = vcl.NewSpinEdit(sf)
	sf.EditMaxConnect.SetMaxValue(1000)
	sf.EditMaxConnect.SetMinValue(1)
	sf.EditMaxConnect.SetHint("单个IP地址,最多可以建立连接数, 超过指定连接数将按下面的操作处理")
	sf.EditMaxConnect.SetShowHint(true)
	sf.EditMaxConnect.SetParentShowHint(false)
	sf.EditMaxConnect.SetValue(50)
	sf.EditMaxConnect.SetBounds(69, 4, 71, 22)

	sf.Label3 = vcl.NewLabel(sf)
	sf.Label3.SetCaption("连接/IP")
	sf.Label3.SetBounds(147, 9, 47, 13)

	sf.GroupBox3 = &TGroupBox3{
		TGroupBox: vcl.NewGroupBox(sf),
	}
	sf.GroupBox3.SetCaption("攻击操作")
	sf.GroupBox3.SetBounds(5, 115, 183, 97)
	sf.GroupBox3.Layout(sender)

	sf.Label7 = vcl.NewLabel(sf)
	sf.Label7.SetCaption("以上参数调后立即生效")
	sf.Label7.Font().SetColor(colors.ClRed)
	sf.Label7.SetBounds(35, 226, 140, 13)

	sf.Label2.SetParent(sf)
	sf.EditMaxConnect.SetParent(sf)
	sf.Label3.SetParent(sf)
	sf.GroupBox3.SetParent(sf)
	sf.Label7.SetParent(sf)
}

func (sf *TGroupBox3) Layout(sender *TFrmIPaddrFilter) {

	sf.RadioDisConnect = vcl.NewRadioButton(sf)
	sf.RadioDisConnect.SetCaption("断开连接")
	sf.RadioDisConnect.SetHint("将连接简单的断开处理")
	sf.RadioDisConnect.SetParentShowHint(false)
	sf.RadioDisConnect.SetShowHint(true)
	sf.RadioDisConnect.SetBounds(17, 5, 140, 19)

	sf.RadioAddTempList = vcl.NewRadioButton(sf)
	sf.RadioAddTempList.SetCaption("加入动态过滤列表")
	sf.RadioAddTempList.SetHint("将此连接的IP加入动态过滤列表, 并将此IP的所有连接强行中断")
	sf.RadioAddTempList.SetParentShowHint(false)
	sf.RadioAddTempList.SetShowHint(true)
	sf.RadioAddTempList.SetBounds(17, 29, 140, 19)

	sf.RadioAddBlockList = vcl.NewRadioButton(sf)
	sf.RadioAddBlockList.SetCaption("加入永久过滤列表")
	sf.RadioAddBlockList.SetHint("将此连接的IP加入永久过滤列表, 并将此IP的所有连接强行中断")
	sf.RadioAddBlockList.SetParentShowHint(false)
	sf.RadioAddBlockList.SetShowHint(true)
	sf.RadioAddBlockList.SetBounds(17, 53, 140, 19)

	sf.RadioDisConnect.SetParent(sf)
	sf.RadioAddTempList.SetParent(sf)
	sf.RadioAddBlockList.SetParent(sf)
}

// ******************** TFrmIPaddrFilter ********************
func (sf *TFrmIPaddrFilter) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.SetBounds(420, 296, 679, 347)
	sf.SetBorderStyle(types.BsSingle)
	sf.Layout()
}
