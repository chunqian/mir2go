// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"bufio"
	"encoding/binary"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// ******************** Enum ********************
// TBlockIPMethod 是处理 IP 阻塞的方法枚举.
type TBlockIPMethod int

const (
	MDisconnect TBlockIPMethod = iota // 断开连接
	MBlock                            // 阻塞
	MBlockList                        // 阻塞列表
)

// ******************** Type ********************
// TSockaddr 存储与 IP 地址相关的信息.
type TSockaddr struct {
	NIPaddr        int
	NCount         int
	DwIPCountTick1 uint32
	NIPCount1      int
	DwIPCountTick2 uint32
	NIPCount2      int
	DwDenyTick     uint32
	NIPDenyCount   int
}

// ******************** Var ********************
var (
	BlockIPList          []TSockaddr // 阻塞 IP 列表
	BlockIPListMutex     sync.Mutex
	BlockMethod          = MDisconnect // 默认的阻塞方法
	BoClose              = false
	BoDecodeLock         bool
	BoGateReady          = false
	BoKeepAliveTimeOut   = false
	BoSendHoldTimeOut    bool
	BoServiceStart       = false
	BoShowMessage        = false
	BoStarted            = false
	CurrIPaddrList       []TSockaddr // 当前 IP 地址列表
	CurrIPaddrListMutex  sync.Mutex
	DwKeepAliveTick      uint32
	DwKeepConnectTimeOut uint32 = 60 * 1000 // 保持连接的超时时间
	DwSendHoldTick       uint32
	GateAddr             = "0.0.0.0"
	GateClass            = "LoginGate"
	GateName             = "登录网关"
	GatePort             = 7000
	GboDynamicIPDisMode = false  // 动态 IP 分发模式
	MainLogMsgList       []string // 存储日志信息的列表
	MainLogMsgListMutex  sync.Mutex
	N456A2C              int
	N456A30              int
	NIPCountLimit1       = 20 // IP 限制次数
	NIPCountLimit2       = 40
	NMaxConnOfIPaddr     = 10 // IP 地址的最大连接数
	NSendMsgCount        int
	NShowLogLevel        = 3              // 显示日志等级
	PosFile              = "./Config.ini" // 配置文件路径
	ServerAddr           = "127.0.0.1"
	ServerPort           = 5500
	StringList456A14     []string
	TempBlockIPList      []TSockaddr // 临时阻塞 IP 列表
	TempBlockIPListMutex sync.Mutex
	TitleName            = "热血传奇"
)

func ipToInteger(ip net.IP) uint32 {
	// 获取 IPv4 部分
	ip = ip.To4()

	// 将 IP 地址转换为 4 字节表示, 然后转换为一个 uint32 整数
	return binary.BigEndian.Uint32(ip)
}

func integerToIP(integer uint32) net.IP {
	// 创建一个 4 字节的 slice
	bytes := make([]byte, 4)

	// 将整数转换为 4 字节表示
	binary.BigEndian.PutUint32(bytes, integer)

	// 将字节转换为 net.IP 类型
	return net.IP(bytes)
}

func LoadBlockIPFile() {
	// 定义文件名
	sFileName := "./BlockIPList.txt"

	// 检查文件是否存在
	if _, err := os.Stat(sFileName); os.IsNotExist(err) {
		return
	}

	// 打开文件
	file, err := os.Open(sFileName)
	if err != nil {
		return
	}
	defer file.Close()

	// 使用 bufio 的 Scanner 逐行读取
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 修整字符串并检查是否为空
		sIPaddr := strings.TrimSpace(scanner.Text())
		if sIPaddr == "" {
			continue
		}

		// 使用 net 包来解析 IP 地址
		ip := net.ParseIP(sIPaddr)
		if ip == nil {
			continue
		}

		// 将 IP 地址加入 BlockIPList
		BlockIPList = append(BlockIPList, TSockaddr{NIPaddr: int(ipToInteger(ip))})
	}

	// 错误处理
	if err := scanner.Err(); err != nil {
		// Handle the error
	}
}

func SaveBlockIPList() {
	// 打开（或创建）文件
	file, err := os.Create("./BlockIPList.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// 使用 bufio 的 Writer 进行写入
	writer := bufio.NewWriter(file)

	for _, sockaddr := range BlockIPList {
		// 使用 integerToIP 自定义函数将整数转换为 net.IP 类型, 再转换为字符串
		ipString := integerToIP(uint32(sockaddr.NIPaddr)).String()

		// 将 IP 写入文件
		writer.WriteString(ipString + "\n")
	}

	// 刷新缓冲区, 确保所有内容都写入文件
	writer.Flush()
}

func GetTickCount() uint32 {
	// 当前时间的毫秒数
	return uint32(time.Now().UnixNano() / 1e6)
}
