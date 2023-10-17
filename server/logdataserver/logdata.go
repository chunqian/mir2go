// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"fmt"
	"os"
	"time"

	. "github.com/chunqian/mir2go/common"
	"github.com/chunqian/mir2go/server/logdataserver/widget"
	log "github.com/chunqian/tinylog"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ******************** Type ********************
type TFrmLogData struct {
	*vcl.TForm

	mainMenu *widget.TMainMenu
	label3   *vcl.TLabel
	memo1    *vcl.TMemo
	socket   *TUdpSocket
	timer1   *vcl.TTimer
}

// ******************** Var ********************
var (
	frmLogData *TFrmLogData
)

// ******************** TFrmLogData ********************
func (sf *TFrmLogData) SetComponents() {

	sf.mainMenu = widget.NewMainMenu(sf)

	sf.label3 = vcl.NewLabel(sf)
	sf.label3.SetCaption("当前日志文件:")
	sf.label3.SetBounds(9, 9, 85, 13)

	sf.memo1 = vcl.NewMemo(sf)
	sf.memo1.SetName("memoLog")
	sf.memo1.SetBounds(11, 30, 303, 75)
	sf.memo1.SetWordWrap(false)
	sf.memo1.SetScrollBars(types.SsHorizontal)
	sf.memo1.SetReadOnly(true)

	sf.label3.SetParent(sf)
	sf.memo1.SetParent(sf)
}

func (sf *TFrmLogData) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.SetCaption("日志服务器")
	sf.EnabledMaximize(false)
	sf.SetBounds(782, 338, 329, 121)

	// constraints := vcl.AsSizeConstraints(sf.Constraints())
	// constraints.SetMaxWidth(500)
	// sf.SetConstraints(constraints)

	sf.SetBorderStyle(types.BsSingle)
	sf.SetComponents()

	sf.timer1 = vcl.NewTimer(sf)
	sf.timer1.SetInterval(3000)
	sf.timer1.SetEnabled(true)
	sf.timer1.SetOnTimer(sf.Timer1Timer)

	RemoteClose = false
	LogMsgList.data = make([]string, 0)

	conf := vcl.NewIniFile(AppDir + "./Config.ini")
	if conf != nil {
		BaseDir = conf.ReadString(ServerClass, "BaseDir", BaseDir)
		ServerName = conf.ReadString(ServerClass, "ServerName", ServerName)
		ServerAddr = conf.ReadString(ServerClass, "LogAddr", ServerAddr)
		ServerPort = conf.ReadInteger(ServerClass, "LogPort", ServerPort)
		conf.Free()
	}

	sf.SetCaption(Caption + " - " + ServerName)
	sf.memo1.SetText(BaseDir)

	// 初始化UDP组件
	sf.socket = &TUdpSocket{}
	sf.socket.ListenUDP(sf, ServerAddr, ServerPort)

	// 注册Observer
	ObserverGetTopic("TFrmLogData").AddObserver(frmLogData)
}

func (sf *TFrmLogData) OnFormDestroy(sender vcl.IObject) {
	LogMsgList.data = LogMsgList.data[:0]
}

func (sf *TFrmLogData) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	*canClose = vcl.MessageDlg("是否确认退出服务器?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.IdYes
}

func (sf *TFrmLogData) UdpSocketError(socket *TUdpSocket, err error) {
	vcl.ShowMessage("读取UDP数据出错: " + err.Error())
}

func (sf *TFrmLogData) UdpSocketRead(socket *TUdpSocket, message string) {
	LogMsgList.data = append(LogMsgList.data, message)
}

func (sf *TFrmLogData) Timer1Timer(sender vcl.IObject) {
	sf.WriteLogFile()
}

func (sf *TFrmLogData) WriteLogFile() {
	LogMsgList.mu.Lock()
	defer LogMsgList.mu.Unlock()

	if len(LogMsgList.data) <= 0 {
		return
	}

	// 获取当前日期和时间
	now := time.Now()
	year, month, day := now.Date()
	hour, min, _ := now.Clock()

	// 构造目录和文件名
	logDir := fmt.Sprintf("%s/%d-%02d-%02d", BaseDir, year, month, day)
	logFile := fmt.Sprintf("%s/Log-%02dh%02dm.txt", logDir, hour, (min/10)*2)

	// 显示文件名
	MainLog.AddLogMsg(logFile, 3)
	sf.showMainLogMsg()

	// 创建目录（如果不存在）
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0755)
	}

	// 打开或创建文件
	fl, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error("Error opening file: {}", err.Error())
		return
	}
	defer fl.Close()

	// 写入日志信息
	for i := 0; i < len(LogMsgList.data); i++ {
		msg := LogMsgList.data[i]
		logEntry := fmt.Sprintf("%s %s\n", msg, now.Format("2006-01-02 15:04:05"))
		fl.WriteString(logEntry)
	}

	// 清空logMsgList
	LogMsgList.data = LogMsgList.data[:0]
}

func (sf *TFrmLogData) showMainLogMsg() {
	// 获取主日志列表, 在 GUI 中显示日志
	memoLog := vcl.AsMemo(sf.FindComponent("memoLog"))

	// 更新UI
	for _, logMsg := range MainLog.MsgList() {
		memoLog.Lines().Add(logMsg)
	}

	// 清空主日志列表
	MainLog.ClearMsgList()
}

func (sf *TFrmLogData) ObserverNotifyReceived(tag string, data interface{}) {
	switch tag {
	case "showMainLogMsg":
		sf.showMainLogMsg()
	case "menuControlClearLogClick":
		sf.memo1.Clear()
		sf.memo1.SetText(BaseDir)
	}
}
