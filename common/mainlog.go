// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package common

import (
	"fmt"
	"sync"
	"time"
)

// ******************** Type ********************
type TMainLog struct {
	msgList  []string   // 存储日志信息的列表
	mutex    sync.Mutex // 列表锁
	logLevel int32      // 显示日志等级
}

// ******************** Var ********************
var (
	MainLog *TMainLog
)

func init() {
	MainLog = &TMainLog{
		msgList:  make([]string, 0),
		mutex:    sync.Mutex{},
		logLevel: 3,
	}
}

func (ml *TMainLog) SetLogLevel(msgLevel int32) {
	ml.logLevel = msgLevel
}

func (ml *TMainLog) AddMsg(msg string) {
	ml.mutex.Lock()
	defer ml.mutex.Unlock()

	ml.msgList = append(ml.msgList, msg)
}

func (ml *TMainLog) AddLogMsg(msg string, msgLevel int32) {
	ml.mutex.Lock()
	defer ml.mutex.Unlock()

	if msgLevel <= ml.logLevel {
		tMsg := fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), msg)
		ml.msgList = append(ml.msgList, tMsg)
	}
}

func (ml *TMainLog) MsgList() []string {
	return ml.msgList
}

func (ml *TMainLog) ClearMsgList() {
	ml.mutex.Lock()
	defer ml.mutex.Unlock()

	ml.msgList = ml.msgList[:0]
}
