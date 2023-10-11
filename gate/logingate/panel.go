// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import "github.com/ying32/govcl/vcl"

type TPanel struct {
	*vcl.TPanel

	Label2 *vcl.TLabel
	Hold   *vcl.TLabel
	Lack   *vcl.TLabel
}

func NewPanel(sender vcl.IComponent) *TPanel {
	sf := &TPanel{
		TPanel: vcl.NewPanel(sender),
	}

	sf.Label2 = vcl.NewLabel(sf)
	sf.Label2.SetCaption("label2")
	sf.Label2.SetBounds(199, 11, 42, 13)

	sf.Lack = vcl.NewLabel(sf)
	sf.Lack.SetCaption("0/0")
	sf.Lack.SetBounds(195, 33, 21, 13)

	sf.Hold = vcl.NewLabel(sf)
	sf.Hold.SetCaption("")
	sf.Hold.SetBounds(106, 10, 7, 13)

	sf.Label2.SetParent(sf)
	sf.Lack.SetParent(sf)
	sf.Hold.SetParent(sf)

	return sf
}
