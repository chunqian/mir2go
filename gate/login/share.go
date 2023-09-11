// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ying32/govcl/vcl"
)

// TBlockIPMethod 是处理 IP 阻塞的方法枚举.
type TBlockIPMethod int

const (
	mDisconnect TBlockIPMethod = iota // 断开连接
	mBlock                            // 阻塞
	mBlockList                        // 阻塞列表
)

// TSockaddr 存储与 IP 地址相关的信息.
type TSockaddr struct {
	nIPaddr        int
	nCount         int
	dwIPCountTick1 uint32
	nIPCount1      int
	dwIPCountTick2 uint32
	nIPCount2      int
	dwDenyTick     uint32
	nIPDenyCount   int
}

var (
	// 全局的互斥锁
	CSMainLog   sync.Mutex
	CSFilterMsg sync.Mutex

	// 存储日志信息的列表
	MainLogMsgList       []string
	mutex_MainLogMsgList sync.Mutex

	// 阻塞 IP 列表
	BlockIPList []TSockaddr
	// 临时阻塞 IP 列表
	TempBlockIPList []TSockaddr
	// 当前 IP 地址列表
	CurrIPaddrList []TSockaddr

	// IP 限制次数
	nIPCountLimit1 = 20
	nIPCountLimit2 = 40

	// 显示日志等级
	nShowLogLevel = 3

	// 其他变量
	StringList456A14 []string

	// 服务器相关设置
	GateClass  = "LoginGate"
	GateName   = "登录网关"
	TitleName  = "热血传奇"
	ServerPort = 5500
	ServerAddr = "127.0.0.1"
	GatePort   = 7000
	GateAddr   = "0.0.0.0"

	// 服务器状态相关
	boGateReady        = false
	boShowMessage      = false
	boStarted          = false
	boClose            = false
	boServiceStart     = false
	dwKeepAliveTick    uint32
	boKeepAliveTimeOut = false
	nSendMsgCount      int
	n456A2C            int
	n456A30            int
	boSendHoldTimeOut  bool
	dwSendHoldTick     uint32
	boDecodeLock       bool

	// IP 地址的最大连接数
	nMaxConnOfIPaddr = 10

	// 默认的阻塞方法
	BlockMethod = mDisconnect

	// 保持连接的超时时间
	dwKeepConnectTimeOut uint32 = 60 * 1000

	// 动态 IP 分发模式
	g_boDynamicIPDisMode = false

	// 启动消息
	g_sNowStartGate = "正在启动登录前置服务器..."
	g_sNowStartOK   = "启动登录前置服务器完成..."

	// 配置文件路径
	PosFile = "./Config.ini"
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
		BlockIPList = append(BlockIPList, TSockaddr{nIPaddr: int(ipToInteger(ip))})
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
		ipString := integerToIP(uint32(sockaddr.nIPaddr)).String()

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

func MainOutMessage(sMsg string, nMsgLevel int) {
	mutex_MainLogMsgList.Lock()
	defer mutex_MainLogMsgList.Unlock()

	if nMsgLevel <= nShowLogLevel {
		tMsg := fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), sMsg)
		MainLogMsgList = append(MainLogMsgList, tMsg)
	}
}

func ShowMainLogMsg(f *TFrmMain) {
	if GetTickCount()-f.dwShowMainLogTick < 200 {
		return
	}
	f.dwShowMainLogTick = GetTickCount()

	f.boShowLocked = true
	defer func() { f.boShowLocked = false }()

	// 获取和清空主日志列表
	mutex_MainLogMsgList.Lock()
	f.tempLogList = append(f.tempLogList, MainLogMsgList...)
	MainLogMsgList = MainLogMsgList[:0]
	mutex_MainLogMsgList.Unlock()

	// 在 GUI 中显示日志
	memoLog := vcl.AsMemo(f.FindComponent("MemoLog"))
	for _, logMsg := range f.tempLogList {
		vcl.ThreadSync(func() {
			memoLog.Lines().Add(logMsg)
		})
	}
	f.tempLogList = f.tempLogList[:0]
}
