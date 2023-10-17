// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	. "github.com/chunqian/mir2go/common"
	"github.com/chunqian/mir2go/gate/logingate/widget"
)

type TFrmGeneralConfig struct {
	*vcl.TForm

	ButtonOK     *vcl.TButton
	GroupBoxNet  *widget.TGroupBoxNet
	GroupBoxInfo *widget.TGroupBoxInfo
}

var (
	frmGeneralConfig *TFrmGeneralConfig
)

// ******************** TFrmGeneralConfig ********************
func (sf *TFrmGeneralConfig) SetComponents() {

	sf.ButtonOK = vcl.NewButton(sf)
	sf.ButtonOK.SetCaption("确定(&O)")
	sf.ButtonOK.SetDefault(true)
	sf.ButtonOK.SetBounds(301, 139, 87, 27)
	sf.ButtonOK.SetParent(sf)

	sf.GroupBoxNet = widget.NewGroupBoxNet(sf)
	sf.GroupBoxNet.SetCaption("网络设置")
	sf.GroupBoxNet.Font().SetSize(9)
	sf.GroupBoxNet.SetBounds(9, 9, 200, 122)

	sf.GroupBoxInfo = widget.NewGroupBoxInfo(sf)
	sf.GroupBoxInfo.SetCaption("基本参数")
	sf.GroupBoxInfo.Font().SetSize(9)
	sf.GroupBoxInfo.SetBounds(217, 9, 174, 122)

	sf.ButtonOK.SetParent(sf)
	sf.GroupBoxNet.SetParent(sf)
	sf.GroupBoxInfo.SetParent(sf)
}

func (sf *TFrmGeneralConfig) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.SetBounds(748, 335, 401, 171)
	sf.SetBorderStyle(types.BsSingle)
	sf.SetComponents()

	// 注册观察者
	ObserverGetTopic("TFrmGeneralConfig").AddObserver(frmGeneralConfig)
}

func (sf *TFrmGeneralConfig) ObserverNotifyReceived(tag string, data interface{}) {
	switch tag {
	case "menuOptionGeneralClick":
		sf.ShowModal()
	}
}
