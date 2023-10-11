// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TGroupBoxNet struct {
	*vcl.TGroupBox

	LabelGateIPAddr   *vcl.TLabel
	LabelGatePort     *vcl.TLabel
	LabelServerIPAddr *vcl.TLabel
	LabelServerPort   *vcl.TLabel

	EditGateIPAddr   *vcl.TEdit
	EditGatePort     *vcl.TEdit
	EditServerIPAddr *vcl.TEdit
	EditServerPort   *vcl.TEdit
}

type TGroupBoxInfo struct {
	*vcl.TGroupBox

	Label1            *vcl.TLabel
	LabelShowLogLevel *vcl.TLabel
	EditTitle         *vcl.TEdit
	TrackBarLogLevel  *vcl.TTrackBar
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

func NewGroupBoxNet(sender vcl.IComponent) *TGroupBoxNet {
	sf := &TGroupBoxNet{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.LabelGateIPAddr = vcl.NewLabel(sf)
	sf.LabelGateIPAddr.SetCaption("网关地址:")
	sf.LabelGateIPAddr.SetBounds(9, 9, 59, 13)

	sf.LabelGatePort = vcl.NewLabel(sf)
	sf.LabelGatePort.SetCaption("网关端口:")
	sf.LabelGatePort.SetBounds(9, 33, 59, 13)

	sf.LabelServerIPAddr = vcl.NewLabel(sf)
	sf.LabelServerIPAddr.SetCaption("服务器地址:")
	sf.LabelServerIPAddr.SetBounds(9, 59, 59, 13)

	sf.LabelServerPort = vcl.NewLabel(sf)
	sf.LabelServerPort.SetCaption("服务器端口:")
	sf.LabelServerPort.SetBounds(9, 85, 59, 13)

	sf.EditGateIPAddr = vcl.NewEdit(sf)
	sf.EditGateIPAddr.SetText("127.0.0.1")
	sf.EditGateIPAddr.SetBounds(87, 5, 105, 20)

	sf.EditGatePort = vcl.NewEdit(sf)
	sf.EditGatePort.SetText("7200")
	sf.EditGatePort.SetBounds(87, 31, 44, 20)

	sf.EditServerIPAddr = vcl.NewEdit(sf)
	sf.EditServerIPAddr.SetText("127.0.0.1")
	sf.EditServerIPAddr.SetBounds(87, 57, 105, 20)

	sf.EditServerPort = vcl.NewEdit(sf)
	sf.EditServerPort.SetText("5000")
	sf.EditServerPort.SetBounds(87, 83, 44, 20)

	sf.LabelGateIPAddr.SetParent(sf)
	sf.LabelGatePort.SetParent(sf)
	sf.LabelServerIPAddr.SetParent(sf)
	sf.LabelServerPort.SetParent(sf)
	sf.EditGateIPAddr.SetParent(sf)
	sf.EditGatePort.SetParent(sf)
	sf.EditServerIPAddr.SetParent(sf)
	sf.EditServerPort.SetParent(sf)

	return sf
}

func NewGroupBoxInfo(sender vcl.IComponent) *TGroupBoxInfo {
	sf := &TGroupBoxInfo{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.Label1 = vcl.NewLabel(sf)
	sf.Label1.SetCaption("标题:")
	sf.Label1.SetBounds(9, 22, 33, 13)

	sf.LabelShowLogLevel = vcl.NewLabel(sf)
	sf.LabelShowLogLevel.SetCaption("显示日志等级:")
	sf.LabelShowLogLevel.SetBounds(9, 48, 85, 13)

	sf.EditTitle = vcl.NewEdit(sf)
	sf.EditTitle.SetText("热血传奇")
	sf.EditTitle.SetBounds(43, 17, 114, 20)

	sf.TrackBarLogLevel = vcl.NewTrackBar(sf)
	sf.TrackBarLogLevel.SetMax(10)
	sf.TrackBarLogLevel.SetMin(1)
	sf.TrackBarLogLevel.SetBounds(9, 61, 157, 27)

	sf.Label1.SetParent(sf)
	sf.LabelShowLogLevel.SetParent(sf)
	sf.EditTitle.SetParent(sf)
	sf.TrackBarLogLevel.SetParent(sf)

	return sf
}

func NewGroupBox1(sender vcl.IComponent) *TGroupBox1 {
	sf := &TGroupBox1{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.LabelTempList = vcl.NewLabel(sf)
	sf.LabelTempList.SetCaption("动态过滤:")
	sf.LabelTempList.SetBounds(9, 9, 59, 13)

	sf.Label1 = vcl.NewLabel(sf)
	sf.Label1.SetCaption("永久过滤:")
	sf.Label1.SetBounds(147, 9, 59, 13)

	sf.ListBoxTempList = NewListBoxTempList(sf)
	sf.ListBoxTempList.SetHint("动态过滤列表, 在此列表中的IP将无法建立连接, 但在程序重新启动时此列表的信息将被清空")
	sf.ListBoxTempList.SetItemHeight(13)
	sf.ListBoxTempList.SetBounds(0, 31, 138, 261)
	sf.ListBoxTempList.SetParentShowHint(false)
	sf.ListBoxTempList.SetShowHint(true)
	sf.ListBoxTempList.SetSorted(true)

	sf.ListBoxBlockList = NewListBoxBlockList(sf)
	sf.ListBoxBlockList.SetHint("永久过滤列表, 在此列表中的IP将无法建立连接, 此列表将保存于配置文件中, 在程序重新启动时会重新加载此列表")
	sf.ListBoxBlockList.SetItemHeight(13)
	sf.ListBoxBlockList.SetBounds(147, 31, 138, 261)
	sf.ListBoxBlockList.SetParentShowHint(false)
	sf.ListBoxBlockList.SetShowHint(true)
	sf.ListBoxBlockList.SetSorted(true)

	sf.LabelTempList.SetParent(sf)
	sf.Label1.SetParent(sf)
	sf.ListBoxTempList.SetParent(sf)
	sf.ListBoxBlockList.SetParent(sf)

	return sf
}

func NewGroupBox2(sender vcl.IComponent) *TGroupBox2 {
	sf := &TGroupBox2{
		TGroupBox: vcl.NewGroupBox(sender),
	}

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

	sf.GroupBox3 = NewGroupBox3(sf)
	sf.GroupBox3.SetCaption("攻击操作")
	sf.GroupBox3.SetBounds(5, 115, 183, 97)

	sf.Label7 = vcl.NewLabel(sf)
	sf.Label7.SetCaption("以上参数调后立即生效")
	sf.Label7.Font().SetColor(colors.ClRed)
	sf.Label7.SetBounds(35, 226, 140, 13)

	sf.Label2.SetParent(sf)
	sf.EditMaxConnect.SetParent(sf)
	sf.Label3.SetParent(sf)
	sf.GroupBox3.SetParent(sf)
	sf.Label7.SetParent(sf)

	return sf
}

func NewGroupBox3(sender vcl.IComponent) *TGroupBox3 {
	sf := &TGroupBox3{
		TGroupBox: vcl.NewGroupBox(sender),
	}

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

	return sf
}

func NewGroupBoxActive(sender vcl.IComponent) *TGroupBoxActive {
	sf := &TGroupBoxActive{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.Label4 = vcl.NewLabel(sf)
	sf.Label4.SetCaption("连接列表:")
	sf.Label4.SetBounds(0, 9, 59, 13)

	sf.ListBoxActiveList = NewListBoxActiveList(sf)
	sf.ListBoxActiveList.SetHint("当前连接的IP地址列表")
	sf.ListBoxActiveList.SetItemHeight(13)
	sf.ListBoxActiveList.SetBounds(0, 31, 138, 261)
	sf.ListBoxActiveList.SetParentShowHint(false)
	sf.ListBoxActiveList.SetShowHint(true)
	sf.ListBoxActiveList.SetSorted(true)

	sf.Label4.SetParent(sf)
	sf.ListBoxActiveList.SetParent(sf)

	return sf
}
