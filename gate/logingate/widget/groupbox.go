// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package widget

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TGroupBoxNet struct {
	*vcl.TGroupBox

	labelGateIPAddr   *vcl.TLabel
	labelGatePort     *vcl.TLabel
	labelServerIPAddr *vcl.TLabel
	labelServerPort   *vcl.TLabel

	editGateIPAddr   *vcl.TEdit
	editGatePort     *vcl.TEdit
	editServerIPAddr *vcl.TEdit
	editServerPort   *vcl.TEdit
}

type TGroupBoxInfo struct {
	*vcl.TGroupBox

	label1            *vcl.TLabel
	labelShowLogLevel *vcl.TLabel
	editTitle         *vcl.TEdit
	trackBarLogLevel  *vcl.TTrackBar
}

type TGroupBox1 struct {
	*vcl.TGroupBox

	labelTempList    *vcl.TLabel
	label1           *vcl.TLabel
	listBoxTempList  *TListBoxTempList
	listBoxBlockList *TListBoxBlockList
}

type TGroupBox2 struct {
	*vcl.TGroupBox

	editMaxConnect *vcl.TSpinEdit
	groupBox3      *TGroupBox3
	label2         *vcl.TLabel
	label3         *vcl.TLabel
	label7         *vcl.TLabel
}

type TGroupBox3 struct {
	*vcl.TGroupBox

	radioAddBlockList *vcl.TRadioButton
	radioAddTempList  *vcl.TRadioButton
	radioDisConnect   *vcl.TRadioButton
}

type TGroupBoxActive struct {
	*vcl.TGroupBox

	label4            *vcl.TLabel
	listBoxActiveList *TListBoxActiveList
}

func NewGroupBoxNet(sender vcl.IComponent) *TGroupBoxNet {
	sf := &TGroupBoxNet{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.labelGateIPAddr = vcl.NewLabel(sf)
	sf.labelGateIPAddr.SetCaption("网关地址:")
	sf.labelGateIPAddr.SetBounds(9, 9, 59, 13)

	sf.labelGatePort = vcl.NewLabel(sf)
	sf.labelGatePort.SetCaption("网关端口:")
	sf.labelGatePort.SetBounds(9, 33, 59, 13)

	sf.labelServerIPAddr = vcl.NewLabel(sf)
	sf.labelServerIPAddr.SetCaption("服务器地址:")
	sf.labelServerIPAddr.SetBounds(9, 59, 59, 13)

	sf.labelServerPort = vcl.NewLabel(sf)
	sf.labelServerPort.SetCaption("服务器端口:")
	sf.labelServerPort.SetBounds(9, 85, 59, 13)

	sf.editGateIPAddr = vcl.NewEdit(sf)
	sf.editGateIPAddr.SetText("127.0.0.1")
	sf.editGateIPAddr.SetBounds(87, 5, 105, 20)

	sf.editGatePort = vcl.NewEdit(sf)
	sf.editGatePort.SetText("7200")
	sf.editGatePort.SetBounds(87, 31, 44, 20)

	sf.editServerIPAddr = vcl.NewEdit(sf)
	sf.editServerIPAddr.SetText("127.0.0.1")
	sf.editServerIPAddr.SetBounds(87, 57, 105, 20)

	sf.editServerPort = vcl.NewEdit(sf)
	sf.editServerPort.SetText("5000")
	sf.editServerPort.SetBounds(87, 83, 44, 20)

	sf.labelGateIPAddr.SetParent(sf)
	sf.labelGatePort.SetParent(sf)
	sf.labelServerIPAddr.SetParent(sf)
	sf.labelServerPort.SetParent(sf)
	sf.editGateIPAddr.SetParent(sf)
	sf.editGatePort.SetParent(sf)
	sf.editServerIPAddr.SetParent(sf)
	sf.editServerPort.SetParent(sf)

	return sf
}

func NewGroupBoxInfo(sender vcl.IComponent) *TGroupBoxInfo {
	sf := &TGroupBoxInfo{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.label1 = vcl.NewLabel(sf)
	sf.label1.SetCaption("标题:")
	sf.label1.SetBounds(9, 22, 33, 13)

	sf.labelShowLogLevel = vcl.NewLabel(sf)
	sf.labelShowLogLevel.SetCaption("显示日志等级:")
	sf.labelShowLogLevel.SetBounds(9, 48, 85, 13)

	sf.editTitle = vcl.NewEdit(sf)
	sf.editTitle.SetText("热血传奇")
	sf.editTitle.SetBounds(43, 17, 114, 20)

	sf.trackBarLogLevel = vcl.NewTrackBar(sf)
	sf.trackBarLogLevel.SetMax(10)
	sf.trackBarLogLevel.SetMin(1)
	sf.trackBarLogLevel.SetBounds(9, 61, 157, 27)

	sf.label1.SetParent(sf)
	sf.labelShowLogLevel.SetParent(sf)
	sf.editTitle.SetParent(sf)
	sf.trackBarLogLevel.SetParent(sf)

	return sf
}

func NewGroupBox1(sender vcl.IComponent) *TGroupBox1 {
	sf := &TGroupBox1{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.labelTempList = vcl.NewLabel(sf)
	sf.labelTempList.SetCaption("动态过滤:")
	sf.labelTempList.SetBounds(9, 9, 59, 13)

	sf.label1 = vcl.NewLabel(sf)
	sf.label1.SetCaption("永久过滤:")
	sf.label1.SetBounds(147, 9, 59, 13)

	sf.listBoxTempList = NewListBoxTempList(sf)
	sf.listBoxTempList.SetHint("动态过滤列表, 在此列表中的IP将无法建立连接, 但在程序重新启动时此列表的信息将被清空")
	sf.listBoxTempList.SetItemHeight(13)
	sf.listBoxTempList.SetBounds(0, 31, 138, 261)
	sf.listBoxTempList.SetParentShowHint(false)
	sf.listBoxTempList.SetShowHint(true)
	sf.listBoxTempList.SetSorted(true)

	sf.listBoxBlockList = NewListBoxBlockList(sf)
	sf.listBoxBlockList.SetHint("永久过滤列表, 在此列表中的IP将无法建立连接, 此列表将保存于配置文件中, 在程序重新启动时会重新加载此列表")
	sf.listBoxBlockList.SetItemHeight(13)
	sf.listBoxBlockList.SetBounds(147, 31, 138, 261)
	sf.listBoxBlockList.SetParentShowHint(false)
	sf.listBoxBlockList.SetShowHint(true)
	sf.listBoxBlockList.SetSorted(true)

	sf.labelTempList.SetParent(sf)
	sf.label1.SetParent(sf)
	sf.listBoxTempList.SetParent(sf)
	sf.listBoxBlockList.SetParent(sf)

	return sf
}

func NewGroupBox2(sender vcl.IComponent) *TGroupBox2 {
	sf := &TGroupBox2{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.label2 = vcl.NewLabel(sf)
	sf.label2.SetCaption("连接限制:")
	sf.label2.SetBounds(9, 9, 59, 13)

	sf.editMaxConnect = vcl.NewSpinEdit(sf)
	sf.editMaxConnect.SetMaxValue(1000)
	sf.editMaxConnect.SetMinValue(1)
	sf.editMaxConnect.SetHint("单个IP地址,最多可以建立连接数, 超过指定连接数将按下面的操作处理")
	sf.editMaxConnect.SetShowHint(true)
	sf.editMaxConnect.SetParentShowHint(false)
	sf.editMaxConnect.SetValue(50)
	sf.editMaxConnect.SetBounds(69, 4, 71, 22)

	sf.label3 = vcl.NewLabel(sf)
	sf.label3.SetCaption("连接/IP")
	sf.label3.SetBounds(147, 9, 47, 13)

	sf.groupBox3 = NewGroupBox3(sf)
	sf.groupBox3.SetCaption("攻击操作")
	sf.groupBox3.SetBounds(5, 115, 183, 97)

	sf.label7 = vcl.NewLabel(sf)
	sf.label7.SetCaption("以上参数调后立即生效")
	sf.label7.Font().SetColor(colors.ClRed)
	sf.label7.SetBounds(35, 226, 140, 13)

	sf.label2.SetParent(sf)
	sf.editMaxConnect.SetParent(sf)
	sf.label3.SetParent(sf)
	sf.groupBox3.SetParent(sf)
	sf.label7.SetParent(sf)

	return sf
}

func NewGroupBox3(sender vcl.IComponent) *TGroupBox3 {
	sf := &TGroupBox3{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.radioDisConnect = vcl.NewRadioButton(sf)
	sf.radioDisConnect.SetCaption("断开连接")
	sf.radioDisConnect.SetHint("将连接简单的断开处理")
	sf.radioDisConnect.SetParentShowHint(false)
	sf.radioDisConnect.SetShowHint(true)
	sf.radioDisConnect.SetBounds(17, 5, 140, 19)

	sf.radioAddTempList = vcl.NewRadioButton(sf)
	sf.radioAddTempList.SetCaption("加入动态过滤列表")
	sf.radioAddTempList.SetHint("将此连接的IP加入动态过滤列表, 并将此IP的所有连接强行中断")
	sf.radioAddTempList.SetParentShowHint(false)
	sf.radioAddTempList.SetShowHint(true)
	sf.radioAddTempList.SetBounds(17, 29, 140, 19)

	sf.radioAddBlockList = vcl.NewRadioButton(sf)
	sf.radioAddBlockList.SetCaption("加入永久过滤列表")
	sf.radioAddBlockList.SetHint("将此连接的IP加入永久过滤列表, 并将此IP的所有连接强行中断")
	sf.radioAddBlockList.SetParentShowHint(false)
	sf.radioAddBlockList.SetShowHint(true)
	sf.radioAddBlockList.SetBounds(17, 53, 140, 19)

	sf.radioDisConnect.SetParent(sf)
	sf.radioAddTempList.SetParent(sf)
	sf.radioAddBlockList.SetParent(sf)

	return sf
}

func NewGroupBoxActive(sender vcl.IComponent) *TGroupBoxActive {
	sf := &TGroupBoxActive{
		TGroupBox: vcl.NewGroupBox(sender),
	}

	sf.label4 = vcl.NewLabel(sf)
	sf.label4.SetCaption("连接列表:")
	sf.label4.SetBounds(0, 9, 59, 13)

	sf.listBoxActiveList = NewListBoxActiveList(sf)
	sf.listBoxActiveList.SetHint("当前连接的IP地址列表")
	sf.listBoxActiveList.SetItemHeight(13)
	sf.listBoxActiveList.SetBounds(0, 31, 138, 261)
	sf.listBoxActiveList.SetParentShowHint(false)
	sf.listBoxActiveList.SetShowHint(true)
	sf.listBoxActiveList.SetSorted(true)

	sf.label4.SetParent(sf)
	sf.listBoxActiveList.SetParent(sf)

	return sf
}
