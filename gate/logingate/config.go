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

	buttonOK     *vcl.TButton
	groupBoxNet  *widget.TGroupBoxNet
	groupBoxInfo *widget.TGroupBoxInfo
}

var (
	frmGeneralConfig *TFrmGeneralConfig
)

// ******************** TFrmGeneralConfig ********************
func (sf *TFrmGeneralConfig) SetComponents() {

	sf.buttonOK = vcl.NewButton(sf)
	sf.buttonOK.SetCaption("确定(&O)")
	sf.buttonOK.SetDefault(true)
	sf.buttonOK.SetBounds(301, 139, 87, 27)
	sf.buttonOK.SetParent(sf)

	sf.groupBoxNet = widget.NewGroupBoxNet(sf)
	sf.groupBoxNet.SetCaption("网络设置")
	sf.groupBoxNet.Font().SetSize(9)
	sf.groupBoxNet.SetBounds(9, 9, 200, 122)

	sf.groupBoxInfo = widget.NewGroupBoxInfo(sf)
	sf.groupBoxInfo.SetCaption("基本参数")
	sf.groupBoxInfo.Font().SetSize(9)
	sf.groupBoxInfo.SetBounds(217, 9, 174, 122)

	sf.buttonOK.SetParent(sf)
	sf.groupBoxNet.SetParent(sf)
	sf.groupBoxInfo.SetParent(sf)
}

func (sf *TFrmGeneralConfig) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.SetBounds(748, 335, 401, 171)
	sf.SetBorderStyle(types.BsSingle)
	sf.SetComponents()

	// 注册Observer
	ObserverGetTopic("TFrmGeneralConfig").AddObserver(frmGeneralConfig)
}

func (sf *TFrmGeneralConfig) ObserverNotifyReceived(tag string, data interface{}) {
	switch tag {
	case "menuOptionGeneralClick":
		sf.ShowModal()
	}
}
