// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TFrmLogData struct {
	*vcl.TForm
	label1 *vcl.TLabel
	timer1 *vcl.TTimer
	memo1  *vcl.TMemo
}

var (
	frmLogData *TFrmLogData
)

var (
	sBaseDir    string = "./LogBase"
	sServerName string = "热血传奇"
	sCaption    string = "引擎日志服务器"
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

	conf := vcl.NewIniFile("./LogData.ini")
	if conf != nil {
		sBaseDir = conf.ReadString("Setup", "BaseDir", sBaseDir)
		sServerName = conf.ReadString("Setup", "Caption", sServerName)
		sServerName = conf.ReadString("Setup", "ServerName", sServerName)
		nServerPort = conf.ReadInteger("Setup", "Port", nServerPort)
		conf.Free()
	}
	f.SetCaption(sCaption + " - " + sServerName)

	f.memo1.SetText(sBaseDir)
}

func (f *TFrmLogData) OnFormCloseQuery(Sender vcl.IObject, CanClose *bool) {
	*CanClose = vcl.MessageDlg("是否确认退出服务器?",
								types.MtConfirmation,
								types.MbYes,
								types.MbNo) == types.IdYes
}

func (f *TFrmLogData) OnTimer1Timer(object vcl.IObject) {
	fmt.Println("OnTimer1Timer.")
}
