//go:build darwin

// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"runtime"

	_ "github.com/chunqian/mir2go/pkgs/macapp"
)

var AppDir string = "/Users/tim/Downloads/github/mir2go/bin/logingate/"

func init() {
	if runtime.GOOS != "darwin" {
		AppDir = ""
	}
}
