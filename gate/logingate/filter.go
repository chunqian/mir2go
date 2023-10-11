// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TFrmIPAddrFilter struct {
	*vcl.TForm

	ButtonOK       *vcl.TButton
	GroupBox1      *TGroupBox1
	GroupBox2      *TGroupBox2
	GroupBoxActive *TGroupBoxActive
}

var (
	FrmIPAddrFilter *TFrmIPAddrFilter
)

// ******************** TFrmIPAddrFilter ********************
func (sf *TFrmIPAddrFilter) SetComponents() {

	sf.ButtonOK = vcl.NewButton(sf)
	sf.ButtonOK.SetCaption("确定(&O)")
	sf.ButtonOK.SetDefault(true)
	sf.ButtonOK.SetBounds(568, 295, 86, 27)

	sf.GroupBoxActive = NewGroupBoxActive(sf)
	sf.GroupBoxActive.SetCaption("当前连接")
	sf.GroupBoxActive.Font().SetSize(9)
	sf.GroupBoxActive.SetBounds(9, 9, 148, 313)

	sf.GroupBox1 = NewGroupBox1(sf)
	sf.GroupBox1.SetCaption("过滤列表")
	sf.GroupBox1.Font().SetSize(9)
	sf.GroupBox1.SetBounds(162, 9, 294, 313)

	sf.GroupBox2 = NewGroupBox2(sf)
	sf.GroupBox2.SetCaption("攻击保护")
	sf.GroupBox2.Font().SetSize(9)
	sf.GroupBox2.SetBounds(464, 9, 201, 274)

	sf.ButtonOK.SetParent(sf)
	sf.GroupBoxActive.SetParent(sf)
	sf.GroupBox1.SetParent(sf)
	sf.GroupBox2.SetParent(sf)
}

func (sf *TFrmIPAddrFilter) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.SetBounds(420, 296, 679, 347)
	sf.SetBorderStyle(types.BsSingle)
	sf.SetComponents()
}
