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

	logMsgList      []string
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

// ******************** Layout ********************
func (sf *TFrmLogData) Layout() {

	sf.Label3 = vcl.NewLabel(sf)
	sf.Label3.SetParent(sf)
	sf.Label3.SetCaption("当前日志文件:")
	sf.Label3.SetBounds(9, 9, 85, 13)

	sf.Timer1 = vcl.NewTimer(sf)
	sf.Timer1.SetInterval(3000)
	sf.Timer1.SetEnabled(true)
	sf.Timer1.SetOnTimer(sf.Timer1Timer)

	sf.Memo1 = vcl.NewMemo(sf)
	sf.Memo1.SetParent(sf)
	sf.Memo1.SetBounds(11, 30, 303, 75)
	sf.Memo1.SetReadOnly(true)
}

// ******************** TFrmLogData ********************
func (sf *TFrmLogData) OnFormCreate(sender vcl.IObject) {

	// 布局
	sf.SetCaption("日志服务器")
	sf.EnabledMaximize(false)
	sf.SetBounds(782, 338, 329, 121)
	// constraints := vcl.AsSizeConstraints(sf.Constraints())
	// constraints.SetMaxWidth(500)
	// sf.SetConstraints(constraints)
	sf.SetBorderStyle(types.BsSingle)
	sf.Layout()

	sf.remoteClose = false

	sf.logMsgList = make([]string, 0)

	conf := vcl.NewIniFile("./Config.ini")
	if conf != nil {
		BaseDir = conf.ReadString("Setup", "BaseDir", BaseDir)
		ServerName = conf.ReadString("Setup", "ServerName", ServerName)
		ServerAddr = conf.ReadString("Setup", "LogAddr", ServerAddr)
		ServerPort = conf.ReadInteger("Setup", "LogPort", ServerPort)
		conf.Free()
	}

	sf.SetCaption(Caption + " - " + ServerName)
	sf.Memo1.SetText(BaseDir)

	// 初始化UDP组件
	println(fmt.Sprintf("%s:%d", ServerAddr, ServerPort))
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ServerAddr, ServerPort))
	if err != nil {
		vcl.ShowMessage("无法解析地址: " + err.Error())
		return
	}

	sf.UdpConn, err = net.ListenUDP("udp", addr)
	if err != nil {
		vcl.ShowMessage("无法监听UDP: " + err.Error())
		return
	}

	// 启动goroutine来接收UDP消息
	go sf.UDPDataReceived()
}

func (sf *TFrmLogData) OnFormDestroy(sender vcl.IObject) {
	sf.logMsgList = sf.logMsgList[:0]
}

func (sf *TFrmLogData) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	*canClose = vcl.MessageDlg("是否确认退出服务器?",
		types.MtConfirmation,
		types.MbYes,
		types.MbNo) == types.IdYes
}

func (sf *TFrmLogData) UDPDataReceived() {
	buffer := make([]byte, 2048) // 最大2048个字节
	for {
		numberBytes, _, err := sf.UdpConn.ReadFromUDP(buffer)
		if err != nil {
			vcl.ThreadSync(func() {
				vcl.ShowMessage("读取UDP数据出错: " + err.Error())
			})
			return
		}
		message := string(buffer[:numberBytes])
		sf.logMsgList = append(sf.logMsgList, message)
	}
}

func (sf *TFrmLogData) Timer1Timer(object vcl.IObject) {
	sf.WriteLogFile()
}

func (sf *TFrmLogData) WriteLogFile() {
	sf.logMsgListMutex.Lock()
	defer sf.logMsgListMutex.Unlock()

	if len(sf.logMsgList) <= 0 {
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
		sf.Memo1.SetText(sLogFile)
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
	for i := 0; i < len(sf.logMsgList); i++ {
		msg := sf.logMsgList[i]
		logEntry := fmt.Sprintf("%s\t%s\n", msg, now.Format("2006-01-02 15:04:05"))
		fl.WriteString(logEntry)
	}

	// 清空logMsgList
	sf.logMsgList = sf.logMsgList[:0]
}
