// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"fmt"
	"os"
	"time"

	. "github.com/chunqian/mir2go/common"
	log "github.com/chunqian/tinylog"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ******************** Type ********************
type TFrmLogData struct {
	*vcl.TForm

	Label3 *vcl.TLabel
	Memo1  *vcl.TMemo
	Socket *TUdpSocket
	Timer1 *vcl.TTimer
}

// ******************** Var ********************
var (
	FrmLogData *TFrmLogData
)

// ******************** TFrmLogData ********************
func (sf *TFrmLogData) SetComponents() {

	sf.Label3 = vcl.NewLabel(sf)
	sf.Label3.SetParent(sf)
	sf.Label3.SetCaption("当前日志文件:")
	sf.Label3.SetBounds(9, 9, 85, 13)

	sf.Memo1 = vcl.NewMemo(sf)
	sf.Memo1.SetParent(sf)
	sf.Memo1.SetBounds(11, 30, 303, 75)
	sf.Memo1.SetReadOnly(true)
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

	sf.Timer1 = vcl.NewTimer(sf)
	sf.Timer1.SetInterval(3000)
	sf.Timer1.SetEnabled(true)
	sf.Timer1.SetOnTimer(sf.Timer1Timer)

	RemoteClose = false
	LogMsgList.Data = make([]string, 0)

	conf := vcl.NewIniFile(AppDir + "./Config.ini")
	if conf != nil {
		BaseDir = conf.ReadString(ServerClass, "BaseDir", BaseDir)
		ServerName = conf.ReadString(ServerClass, "ServerName", ServerName)
		ServerAddr = conf.ReadString(ServerClass, "LogAddr", ServerAddr)
		ServerPort = conf.ReadInteger(ServerClass, "LogPort", ServerPort)
		conf.Free()
	}

	sf.SetCaption(Caption + " - " + ServerName)
	sf.Memo1.SetText(BaseDir)

	// 初始化UDP组件
	sf.Socket = &TUdpSocket{}
	sf.Socket.ListenUDP(sf, ServerAddr, ServerPort)
}

func (sf *TFrmLogData) OnFormDestroy(sender vcl.IObject) {
	LogMsgList.Data = LogMsgList.Data[:0]
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
	LogMsgList.Data = append(LogMsgList.Data, message)
}

func (sf *TFrmLogData) Timer1Timer(sender vcl.IObject) {
	sf.WriteLogFile()
}

func (sf *TFrmLogData) WriteLogFile() {
	LogMsgList.Mu.Lock()
	defer LogMsgList.Mu.Unlock()

	if len(LogMsgList.Data) <= 0 {
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
	sf.Memo1.SetText(logFile)

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
	for i := 0; i < len(LogMsgList.Data); i++ {
		msg := LogMsgList.Data[i]
		logEntry := fmt.Sprintf("%s %s\n", msg, now.Format("2006-01-02 15:04:05"))
		fl.WriteString(logEntry)
	}

	// 清空logMsgList
	LogMsgList.Data = LogMsgList.Data[:0]
}
