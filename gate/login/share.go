// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"bufio"
	"net"
	"os"
	"strings"
	"sync"

	common "github.com/chunqian/mir2go/common"
)

// ******************** Var ********************
var (
	BlockIPList         []common.TSockAddr  // 阻塞 IP 列表
	BlockMethod         = common.Disconnect // 默认的阻塞方法
	DecodeLock          bool
	GateReady           = false
	KeepAliveTimeOut    = false
	SendHoldTimeOut     bool
	ServiceStart        = false
	Started             = false
	Closed              = false
	CurrIPaddrList      []common.TSockAddr // 当前 IP 地址列表
	KeepAliveTick       uint32
	KeepConnectTimeOut  int32 = 60 * 1000 // 保持连接的超时时间
	SendHoldTick        uint32
	GateAddr                     = "0.0.0.0"
	GateClass                    = "Setup"
	GateName                     = "登录网关"
	GatePort            int32    = 7000
	DynamicIPDisMode             = false // 动态 IP 分发模式
	MainLogMsgList      []string         // 存储日志信息的列表
	MainLogMsgListMutex sync.Mutex
	TotalMsgListCount   int
	ActiveConnections   int
	IPCountLimit1             = 20 // IP 限制次数
	IPCountLimit2             = 40
	MaxConnOfIPaddr     int32 = 10 // IP 地址的最大连接数
	SendMsgCount        int
	ShowLogLevel        int32              = 3                       // 显示日志等级
	ConfigFile                             = AppDir + "./Config.ini" // 配置文件路径
	ServerAddr                             = "127.0.0.1"
	ServerPort          int32              = 5500
	TempBlockIPList     []common.TSockAddr // 临时阻塞 IP 列表
	TitleName           = "热血传奇"
)

func LoadBlockIPFile() {
	// 定义文件名
	sFileName := AppDir + "./BlockIPList.txt"

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
		BlockIPList = append(BlockIPList, common.TSockAddr{IPaddr: int(common.IpToInteger(ip))})
	}

	// 错误处理
	if err := scanner.Err(); err != nil {
		// Handle the error
	}
}

func SaveBlockIPList() {
	// 打开（或创建）文件
	file, err := os.Create(AppDir + "./BlockIPList.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// 使用 bufio 的 Writer 进行写入
	writer := bufio.NewWriter(file)

	for _, sockaddr := range BlockIPList {
		// 使用 integerToIP 自定义函数将整数转换为 net.IP 类型, 再转换为字符串
		ipString := common.IntegerToIP(uint32(sockaddr.IPaddr)).String()

		// 将 IP 写入文件
		writer.WriteString(ipString + "\n")
	}

	// 刷新缓冲区, 确保所有内容都写入文件
	writer.Flush()
}
