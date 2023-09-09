// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.DEBUG = true
	vcl.RunApp(&frmMain)
}
