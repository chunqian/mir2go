// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	. "github.com/chunqian/mir2go/common"
)

// ******************** Const ********************
const (
	GATEMAXSESSION   = 1000
	MSGMAXLENGTH     = 20000
	SENDCHECKSIZE    = 512
	SENDCHECKSIZEMAX = 2048
)

// ******************** Type ********************
type TSessionInfo struct {
	socket           *TClientSocket
	socData          string
	sendData         string
	userListIndex    int
	packetIdx        int
	packetErrCount   int
	startLogon       bool
	sendLock         bool
	overNomSize      bool
	overNomSizeCount int8
	sendLatestTime   uint32
	checkSendLength  int
	sendAvailable    bool
	sendCheck        bool
	timeOutTime      uint32
	receiveLength    int
	receiveTick      uint32
	sckHandle        int
	remoteAddr       string
	sayMsgTick       uint32
	hitTick          uint32
}

type TSendUserData struct {
	socketIdx    int
	socketHandle int
	msg          string
}

// ******************** Var ********************
var (
	MainLogMsgList         []string
	ShowLogLevel           int32    = 3
	GateClass                       = "Setup"
	GateName                        = "游戏网关"
	TitleName                       = "热血传奇"
	ServerAddr                      = "127.0.0.1"
	ServerPort             int32    = 5000
	GateAddr                        = "0.0.0.0"
	GatePort               int32    = 7200
	Started                         = false
	Closed                          = false
	ShowBite                        = true // 显示B 或 KB
	ServiceStart                    = true
	GateReady                       = true  // 网关是否就绪
	CheckServerFail                 = false // 网关 <-> 游戏服务器之间检测是否失败 (超时)
	CheckServerTimeOutTime uint32   = 3 * 60 * 1000
	WordFilterList         []string // 文字过滤列表
	SessionArray           [GATEMAXSESSION]*TSessionInfo
	SessionCount           int32
	ShowSckData            bool

	ReplaceWord      = "*"
	ReviceMsgList    []string
	SendMsgList      []string
	CurrConnCount    int32
	SendHoldTimeOut  bool
	SendHoldTick     uint32
	n45AA80          int32
	n45AA84          int32
	CheckRecviceTick uint32
	CheckRecviceMin  uint32
	CheckRecviceMax  uint32

	CheckServerTick           uint32
	CheckServerTimeMin        uint32
	CheckServerTimeMax        uint32
	SocketBuffer              *byte
	BuffLen                   int32
	List_45AA58               []string
	DecodeMsgLock             bool
	ProcessReviceMsgTimeLimit uint32
	ProcessSendMsgTimeLimit   uint32
	BlockIPList               []string // 禁止连接IP列表
	TempBlockIPList           []string // 临时禁止连接IP列表
	MaxConnOfIPaddr           int32    = 50
	MaxClientPacketSize       int32    = 7000
	NomClientPacketSize       int32    = 150
	ClientCheckTimeOut        uint32   = 50
	MaxOverNomSizeCount       int32    = 2
	MaxClientMsgCount         int32    = 15
	BlockMethod                        = Disconnect
	kickOverPacketSize                 = true

	ClientSendBlockSize int32  = 1000 // 发送给客户端数据包大小限制
	ClientTimeOutTime   uint32 = 5000 // 客户端连接会话超时(指定时间内未有数据传输)
	ConfigFileName      string = "./Config.ini"
	SayMsgMaxLen        int32  = 70   // 发言字符长度
	SayMsgTime          uint32 = 1000 // 发主间隔时间
	HitTime             uint32 = 300  // 攻击间隔时间
	SessionTimeOutTime  uint32 = 60 * 60 * 1000

	ShowMainLogTick   uint32
	ShowLocked        bool
	TempLogList       []string
	CheckClientTick   uint32
	ProcessPacketTick uint32

	ServerReady          bool
	LoopCheckTick        uint32
	LoopTime             uint32
	ProcessServerMsgTime uint32
	ProcessClientMsgTime uint32
	ReConnectServerTime  uint32
	RefConsolMsgTick     uint32
	BufferOfM2Size       int32
	RefConsoleMsgTick    uint32
	ReviceMsgSize        int32
	DeCodeMsgSize        int32
	SendBlockSize        int32
	ProcessMsgSize       int32
	HumLogonMsgSize      int32
	HumPlayMsgSize       int32
)
