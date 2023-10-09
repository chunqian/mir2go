// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"sync"
)

// ******************** Var ********************
var (
	BaseDir     string = AppDir + "./LogBase"
	ServerClass string = "Setup"
	ServerName  string = "热血传奇"
	Caption     string = "引擎日志服务器"
	ServerPort  int32  = 10000
	ServerAddr  string = "127.0.0.1"

	LogMsgList      []string
	LogMsgListMutex sync.Mutex
	RemoteClose     bool
)
