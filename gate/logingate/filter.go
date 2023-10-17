// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	. "github.com/chunqian/mir2go/common"
	"github.com/chunqian/mir2go/gate/logingate/widget"
)

type TFrmIPAddrFilter struct {
	*vcl.TForm

	buttonOK       *vcl.TButton
	groupBox1      *widget.TGroupBox1
	groupBox2      *widget.TGroupBox2
	groupBoxActive *widget.TGroupBoxActive
}

var (
	frmIPAddrFilter *TFrmIPAddrFilter
)

// ******************** TFrmIPAddrFilter ********************
func (sf *TFrmIPAddrFilter) SetComponents() {

	sf.buttonOK = vcl.NewButton(sf)
	sf.buttonOK.SetCaption("确定(&O)")
	sf.buttonOK.SetDefault(true)
	sf.buttonOK.SetBounds(568, 295, 86, 27)

	sf.groupBoxActive = widget.NewGroupBoxActive(sf)
	sf.groupBoxActive.SetCaption("当前连接")
	sf.groupBoxActive.Font().SetSize(9)
	sf.groupBoxActive.SetBounds(9, 9, 148, 313)

	sf.groupBox1 = widget.NewGroupBox1(sf)
	sf.groupBox1.SetCaption("过滤列表")
	sf.groupBox1.Font().SetSize(9)
	sf.groupBox1.SetBounds(162, 9, 294, 313)

	sf.groupBox2 = widget.NewGroupBox2(sf)
	sf.groupBox2.SetCaption("攻击保护")
	sf.groupBox2.Font().SetSize(9)
	sf.groupBox2.SetBounds(464, 9, 201, 274)

	sf.buttonOK.SetParent(sf)
	sf.groupBoxActive.SetParent(sf)
	sf.groupBox1.SetParent(sf)
	sf.groupBox2.SetParent(sf)
}

func (sf *TFrmIPAddrFilter) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.SetBounds(420, 296, 679, 347)
	sf.SetBorderStyle(types.BsSingle)
	sf.SetComponents()

	// 注册Observer
	ObserverGetTopic("TFrmIPAddrFilter").AddObserver(frmIPAddrFilter)
}

func (sf *TFrmIPAddrFilter) ObserverNotifyReceived(tag string, data interface{}) {
	switch tag {
	case "menuOptionIpFilterClick":
		sf.ShowModal()
	}
}
