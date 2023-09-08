// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TFrmLogData struct {
	*vcl.TForm
	udpConn *net.UDPConn
	label1  *vcl.TLabel
	timer1  *vcl.TTimer
	memo1   *vcl.TMemo

	logMsgList *vcl.TStringList
	mu         sync.Mutex

	remoteClose bool
}

var (
	frmLogData *TFrmLogData
)

var (
	sBaseDir    string = "./LogBase"
	sServerName string = "热血传奇"
	sCaption    string = "引擎日志服务器"
	nServerAddr string = "127.0.0.1"
	nServerPort int32  = 10000
)

// ******************** TFrmLogData ********************
func (f *TFrmLogData) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("日志服务器")
	f.EnabledMaximize(false)
	f.SetWidth(329)
	f.SetHeight(121)
	f.SetTop(338)
	f.SetLeft(782)

	f.label1 = vcl.NewLabel(f)
	f.label1.SetParent(f)
	f.label1.SetCaption("当前日志文件:")
	f.label1.SetTop(9)
	f.label1.SetLeft(9)
	f.label1.SetHeight(13)
	f.label1.SetWidth(85)

	f.timer1 = vcl.NewTimer(f)
	f.timer1.SetInterval(3000)
	f.timer1.SetEnabled(true)
	f.timer1.SetOnTimer(f.OnTimer1Timer)

	f.memo1 = vcl.NewMemo(f)
	f.memo1.SetParent(f)
	f.memo1.SetTop(30)
	f.memo1.SetLeft(11)
	f.memo1.SetHeight(75)
	f.memo1.SetWidth(303)
	f.memo1.SetReadOnly(true)

	constraints := vcl.AsSizeConstraints(f.Constraints())
	constraints.SetMaxWidth(500)
	f.SetConstraints(constraints)

	f.remoteClose = false

	f.logMsgList = vcl.NewStringList()

	conf := vcl.NewIniFile("./LogData.ini")
	if conf != nil {
		sBaseDir = conf.ReadString("Setup", "BaseDir", sBaseDir)
		sServerName = conf.ReadString("Setup", "Caption", sServerName)
		sServerName = conf.ReadString("Setup", "ServerName", sServerName)
		nServerAddr = conf.ReadString("Setup", "LogAddr", nServerAddr)
		nServerPort = conf.ReadInteger("Setup", "LogPort", nServerPort)
		conf.Free()
	}
	f.SetCaption(sCaption + " - " + sServerName)

	f.memo1.SetText(sBaseDir)

	// 初始化UDP组件
	println(fmt.Sprintf("%s:%d", nServerAddr, nServerPort))
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", nServerAddr, nServerPort))
	if err != nil {
		vcl.ShowMessage("无法解析地址: " + err.Error())
		return
	}

	f.udpConn, err = net.ListenUDP("udp", addr)
	if err != nil {
		vcl.ShowMessage("无法监听UDP: " + err.Error())
		return
	}

	// 启动一个goroutine来接收UDP消息
	go f.dataReceived()
}

func (f *TFrmLogData) OnFormDestroy(Sender vcl.IObject) {
	f.logMsgList.Free()
}

func (f *TFrmLogData) OnFormCloseQuery(Sender vcl.IObject, CanClose *bool) {
	*CanClose = vcl.MessageDlg("是否确认退出服务器?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.IdYes
}

func (f *TFrmLogData) OnTimer1Timer(object vcl.IObject) {
	f.WriteLogFile()
}

func (f *TFrmLogData) dataReceived() {
	buffer := make([]byte, 2048) // 最大2048个字节
	for {
		numberBytes, _, err := f.udpConn.ReadFromUDP(buffer)
		if err != nil {
			vcl.ThreadSync(func() {
				vcl.ShowMessage("读取UDP数据出错: " + err.Error())
			})
			return
		}
		message := string(buffer[:numberBytes])
		f.logMsgList.Add(message)
	}
}

func (f *TFrmLogData) WriteLogFile() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.logMsgList.Count() <= 0 {
		return
	}

	// 获取当前日期和时间
	now := time.Now()
	year, month, day := now.Date()
	hour, min, _ := now.Clock()

	// 构造目录和文件名
	sLogDir := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	sLogFile := fmt.Sprintf("%s/Log-%02dh%02dm.txt", sLogDir, hour, (min/10)*2)

	// 显示文件名
	vcl.ThreadSync(func() {
		f.memo1.SetText(sLogFile)
	})

	// 创建目录（如果不存在）
	if _, err := os.Stat(sLogDir); os.IsNotExist(err) {
		os.Mkdir(sLogDir, 0755)
	}

	// 打开或创建文件
	fl, err := os.OpenFile(sLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fl.Close()

	// 写入日志信息
	for i := int32(0); i < f.logMsgList.Count(); i++ {
		msg := f.logMsgList.Strings(i)
		logEntry := fmt.Sprintf("%s\t%s\n", msg, now.Format("2006-01-02 15:04:05"))
		fl.WriteString(logEntry)
	}

	// 清空LogMsgList
	f.logMsgList.Clear()
}
