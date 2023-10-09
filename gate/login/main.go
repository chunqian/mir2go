// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
)

var (
	AppDir string = "/Users/tim/Downloads/github/mir2go/bin/logingate/"
)

func main() {
	vcl.DEBUG = true
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&FrmMain)
	vcl.Application.CreateForm(&FrmIPAddrFilter)
	vcl.Application.CreateForm(&FrmGeneralConfig)
	vcl.Application.Run()
}
