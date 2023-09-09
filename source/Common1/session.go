// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package common

import (
	"net"
	"sync"

	"github.com/ying32/govcl/vcl"
)

const GATEMAXSESSION = 10000

type TUserSession struct {
	Socket            *net.Conn        // 0x00 可以根据需要使用不同类型
	SRemoteIPaddr     string           // 0x04 远程IP地址
	NSendMsgLen       int              // 0x08 发送消息长度
	Bo0C              bool             // 0x0C 布尔值
	Dw10Tick          uint32           // 0x10 时间戳
	NCheckSendLength  int              // 0x14 检查发送长度
	BoSendAvailable   bool             // 0x18 是否可发送
	BoSendCheck       bool             // 0x19 发送检查
	DwSendLockTimeOut uint32           // 0x1C 发送锁超时
	N20               int              // 0x20 整数字段（具体含义取决于上下文）
	DwUserTimeOutTick uint32           // 0x24 用户超时检查
	SocketHandle      int              // 0x28 套接字句柄（如果你有net.Conn, 这可能是多余的）
	SIP               string           // 0x2C IP地址
	MsgList           *vcl.TStringList // 0x30 消息列表
	DwConnctCheckTick uint32           // 连接数据传输空闲超时检测
	Mutex             sync.Mutex       // 用于并发场景下安全访问MsgList
}

type TSessionArray [GATEMAXSESSION]TUserSession
