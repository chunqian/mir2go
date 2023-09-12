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

// ******************** Type ********************
type TFrmLogData struct {
	*vcl.TForm

	logMsgList      *vcl.TStringList
	logMsgListMutex sync.Mutex
	remoteClose     bool

	Label3  *vcl.TLabel
	Memo1   *vcl.TMemo
	UdpConn *net.UDPConn
	Timer1  *vcl.TTimer
}

// ******************** Var ********************
var (
	FrmLogData *TFrmLogData
)

var (
	BaseDir    string = "./LogBase"
	ServerName string = "热血传奇"
	Caption    string = "引擎日志服务器"
	ServerPort int32  = 10000
	ServerAddr string = "127.0.0.1"
)

// ******************** TFrmLogData ********************
func (f *TFrmLogData) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("日志服务器")
	f.EnabledMaximize(false)
	f.SetWidth(329)
	f.SetHeight(121)
	f.SetTop(338)
	f.SetLeft(782)

	f.Label3 = vcl.NewLabel(f)
	f.Label3.SetParent(f)
	f.Label3.SetCaption("当前日志文件:")
	f.Label3.SetTop(9)
	f.Label3.SetLeft(9)
	f.Label3.SetHeight(13)
	f.Label3.SetWidth(85)

	f.Timer1 = vcl.NewTimer(f)
	f.Timer1.SetInterval(3000)
	f.Timer1.SetEnabled(true)
	f.Timer1.SetOnTimer(f.OnTimer1Timer)

	f.Memo1 = vcl.NewMemo(f)
	f.Memo1.SetParent(f)
	f.Memo1.SetTop(30)
	f.Memo1.SetLeft(11)
	f.Memo1.SetHeight(75)
	f.Memo1.SetWidth(303)
	f.Memo1.SetReadOnly(true)

	constraints := vcl.AsSizeConstraints(f.Constraints())
	constraints.SetMaxWidth(500)
	f.SetConstraints(constraints)

	f.remoteClose = false

	f.logMsgList = vcl.NewStringList()

	conf := vcl.NewIniFile("./Config.ini")
	if conf != nil {
		BaseDir = conf.ReadString("Setup", "BaseDir", BaseDir)
		ServerName = conf.ReadString("Setup", "ServerName", ServerName)
		ServerAddr = conf.ReadString("Setup", "LogAddr", ServerAddr)
		ServerPort = conf.ReadInteger("Setup", "LogPort", ServerPort)
		conf.Free()
	}
	f.SetCaption(Caption + " - " + ServerName)

	f.Memo1.SetText(BaseDir)

	// 初始化UDP组件
	println(fmt.Sprintf("%s:%d", ServerAddr, ServerPort))
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ServerAddr, ServerPort))
	if err != nil {
		vcl.ShowMessage("无法解析地址: " + err.Error())
		return
	}

	f.UdpConn, err = net.ListenUDP("udp", addr)
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
	f.writeLogFile()
}

func (f *TFrmLogData) dataReceived() {
	buffer := make([]byte, 2048) // 最大2048个字节
	for {
		numberBytes, _, err := f.UdpConn.ReadFromUDP(buffer)
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

func (f *TFrmLogData) writeLogFile() {
	f.logMsgListMutex.Lock()
	defer f.logMsgListMutex.Unlock()

	if f.logMsgList.Count() <= 0 {
		return
	}

	// 获取当前日期和时间
	now := time.Now()
	year, month, day := now.Date()
	hour, min, _ := now.Clock()

	// 构造目录和文件名
	sLogDir := fmt.Sprintf("%s/%d-%02d-%02d", BaseDir, year, month, day)
	sLogFile := fmt.Sprintf("%s/Log-%02dh%02dm.txt", sLogDir, hour, (min/10)*2)

	// 显示文件名
	vcl.ThreadSync(func() {
		f.Memo1.SetText(sLogFile)
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

	// 清空logMsgList
	f.logMsgList.Clear()
}
