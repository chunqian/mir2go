// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TFrmGeneralConfig struct {
	*vcl.TForm

	ButtonOK     *vcl.TButton
	GroupBoxNet  *TGroupBoxNet
	GroupBoxInfo *TGroupBoxInfo
}

type TGroupBoxNet struct {
	*vcl.TGroupBox

	LabelGateIPaddr   *vcl.TLabel
	LabelGatePort     *vcl.TLabel
	LabelServerIPaddr *vcl.TLabel
	LabelServerPort   *vcl.TLabel

	EditGateIPaddr   *vcl.TEdit
	EditGatePort     *vcl.TEdit
	EditServerIPaddr *vcl.TEdit
	EditServerPort   *vcl.TEdit
}

type TGroupBoxInfo struct {
	*vcl.TGroupBox

	Label1            *vcl.TLabel
	LabelShowLogLevel *vcl.TLabel
	EditTitle         *vcl.TEdit
	TrackBarLogLevel  *vcl.TTrackBar
}

var (
	FrmGeneralConfig *TFrmGeneralConfig
)

// ******************** Layout ********************
func (sf *TFrmGeneralConfig) Layout() {

	sf.ButtonOK = vcl.NewButton(sf)
	sf.ButtonOK.SetCaption("确定(&O)")
	sf.ButtonOK.SetDefault(true)
	sf.ButtonOK.SetBounds(301, 139, 87, 27)
	sf.ButtonOK.SetParent(sf)

	sf.GroupBoxNet = &TGroupBoxNet{
		TGroupBox: vcl.NewGroupBox(sf),
	}
	sf.GroupBoxNet.SetCaption("网络设置")
	sf.GroupBoxNet.Font().SetSize(9)
	sf.GroupBoxNet.SetBounds(9, 9, 200, 122)
	sf.GroupBoxNet.Layout(sf)

	sf.GroupBoxInfo = &TGroupBoxInfo{
		TGroupBox: vcl.NewGroupBox(sf),
	}
	sf.GroupBoxInfo.SetCaption("基本参数")
	sf.GroupBoxInfo.Font().SetSize(9)
	sf.GroupBoxInfo.SetBounds(217, 9, 174, 122)
	sf.GroupBoxInfo.Layout(sf)

	sf.ButtonOK.SetParent(sf)
	sf.GroupBoxNet.SetParent(sf)
	sf.GroupBoxInfo.SetParent(sf)
}

func (sf *TGroupBoxNet) Layout(sender *TFrmGeneralConfig) {

	sf.LabelGateIPaddr = vcl.NewLabel(sf)
	sf.LabelGateIPaddr.SetCaption("网关地址:")
	sf.LabelGateIPaddr.SetBounds(9, 9, 59, 13)

	sf.LabelGatePort = vcl.NewLabel(sf)
	sf.LabelGatePort.SetCaption("网关端口:")
	sf.LabelGatePort.SetBounds(9, 33, 59, 13)

	sf.LabelServerIPaddr = vcl.NewLabel(sf)
	sf.LabelServerIPaddr.SetCaption("服务器地址:")
	sf.LabelServerIPaddr.SetBounds(9, 59, 59, 13)

	sf.LabelServerPort = vcl.NewLabel(sf)
	sf.LabelServerPort.SetCaption("服务器端口:")
	sf.LabelServerPort.SetBounds(9, 85, 59, 13)

	sf.EditGateIPaddr = vcl.NewEdit(sf)
	sf.EditGateIPaddr.SetText("127.0.0.1")
	sf.EditGateIPaddr.SetBounds(87, 5, 105, 20)

	sf.EditGatePort = vcl.NewEdit(sf)
	sf.EditGatePort.SetText("7200")
	sf.EditGatePort.SetBounds(87, 31, 44, 20)

	sf.EditServerIPaddr = vcl.NewEdit(sf)
	sf.EditServerIPaddr.SetText("127.0.0.1")
	sf.EditServerIPaddr.SetBounds(87, 57, 105, 20)

	sf.EditServerPort = vcl.NewEdit(sf)
	sf.EditServerPort.SetText("5000")
	sf.EditServerPort.SetBounds(87, 83, 44, 20)

	sf.LabelGateIPaddr.SetParent(sf)
	sf.LabelGatePort.SetParent(sf)
	sf.LabelServerIPaddr.SetParent(sf)
	sf.LabelServerPort.SetParent(sf)
	sf.EditGateIPaddr.SetParent(sf)
	sf.EditGatePort.SetParent(sf)
	sf.EditServerIPaddr.SetParent(sf)
	sf.EditServerPort.SetParent(sf)
}

func (sf *TGroupBoxInfo) Layout(sender *TFrmGeneralConfig) {

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
}

// ******************** TFrmGeneralConfig ********************
func (sf *TFrmGeneralConfig) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.SetBounds(748, 335, 401, 171)
	sf.SetBorderStyle(types.BsSingle)
	sf.Layout()
}
