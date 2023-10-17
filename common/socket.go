// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
	"unsafe"

	log "github.com/chunqian/tinylog"
	"github.com/ying32/govcl/vcl"
)

// ******************** Enum ********************
// TBlockIPMethod 是处理 IP 阻塞的方法枚举.
type TBlockIPMethod int32

const (
	Disconnect TBlockIPMethod = iota // 断开连接
	Block                            // 阻塞
	BlockList                        // 阻塞列表
)

// ******************** Type ********************
// TSockAddr 存储与 IP 地址相关的信息.
type TSockAddr struct {
	IP           int
	Count        int
	IPCountTick1 uint32
	IPCount1     int
	IPCountTick2 uint32
	IPCount2     int
	DenyTick     uint32
	IPDenyCount  int
}

type TClientSocket struct {
	*net.TCPConn

	socketHandle uintptr
	Index        int
}

type TServerSocket struct {
	*net.TCPListener

	active            bool
	activeConnections int
	connections       []*TClientSocket
}

type TUdpSocket struct {
	*net.UDPConn
}

// ******************** Interface ********************
type IServerSocket interface {
	ServerSocketClientConnect(s *TClientSocket)
	ServerSocketClientDisconnect(conn *TClientSocket, err error)
	ServerSocketClientError(conn *TClientSocket, err error)
	ServerSocketClientRead(conn *TClientSocket, message string)
}

type IClientSocket interface {
	ClientSocketConnect(socket *TClientSocket)
	ClientSocketDisconnect(socket *TClientSocket, err error)
	ClientSocketError(socket *TClientSocket, err error)
	ClientSocketRead(socket *TClientSocket, message string)
}

type IUdpSocket interface {
	UdpSocketError(socket *TUdpSocket, err error)
	UdpSocketRead(socket *TUdpSocket, message string)
}

func IpToInteger(ip net.IP) uint32 {
	// 获取 IPv4 部分
	ip = ip.To4()

	// 将 IP 地址转换为 4 字节表示, 然后转换为一个 uint32 整数
	return binary.BigEndian.Uint32(ip)
}

func IntegerToIP(integer uint32) net.IP {
	// 创建一个 4 字节的 slice
	bytes := make([]byte, 4)

	// 将整数转换为 4 字节表示
	binary.BigEndian.PutUint32(bytes, integer)

	// 将字节转换为 net.IP 类型
	return net.IP(bytes)
}

func GetAddrHost(n net.Addr) string {
	addr := n.String()
	host, _, _ := net.SplitHostPort(addr)
	return host
}

func GetAddrPort(n net.Addr) string {
	addr := n.String()
	_, port, _ := net.SplitHostPort(addr)
	return port
}

func GetTickCount() uint32 {
	// 当前时间的毫秒数
	return uint32(time.Now().UnixNano() / 1e6)
}

func (s *TServerSocket) messageProducer(iSocket IServerSocket, conn *TClientSocket) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	var dataBuffer bytes.Buffer
	reading := false

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			for i := 0; i < s.activeConnections; {
				socketHandle := s.connections[i].SocketHandle()
				if conn.socketHandle == socketHandle {
					s.connections = append(s.connections[:i], s.connections[i+1:]...)
					break
				} else {
					i++
				}
			}
			s.activeConnections--

			if err != io.EOF {
				vcl.ThreadSync(func() {
					iSocket.ServerSocketClientError(conn, err)
				})
			} else {
				conn.Close()
				vcl.ThreadSync(func() {
					iSocket.ServerSocketClientDisconnect(conn, err)
				})
			}
			break
		}

		for i := 0; i < n; i++ {
			if buffer[i] == '%' {
				reading = true
				dataBuffer.Reset()
				continue
			}

			if buffer[i] == '$' {
				reading = false
				message := dataBuffer.String()
				vcl.ThreadSync(func() {
					iSocket.ServerSocketClientRead(conn, message)
				})
				dataBuffer.Reset()
				continue
			}

			if reading {
				dataBuffer.WriteByte(buffer[i])
			}
		}
	}
}

func (s *TServerSocket) socketProducer(iSocket IServerSocket, ch chan *TClientSocket) {
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Error(err.Error())
			s.active = false
			return
		}
		tcpConn, ok := conn.(*net.TCPConn)
		if !ok {
			log.Error("Not a TCP connection")
			return
		}
		clientSocket := &TClientSocket{
			TCPConn:      tcpConn,
			socketHandle: uintptr(unsafe.Pointer(tcpConn)),
		}

		s.activeConnections++
		s.connections = append(s.connections, clientSocket)
		vcl.ThreadSync(func() {
			iSocket.ServerSocketClientConnect(clientSocket)
		})

		ch <- clientSocket
	}
}

func (s *TServerSocket) socketConsumer(iSocket IServerSocket, ch chan *TClientSocket) {
	for {
		select {
		case sock := <-ch:
			go s.messageProducer(iSocket, sock)
			// default:
			// 	log.Error("Could not read from socketChan")
		}
	}
}

func (s *TServerSocket) Listen(iSocket IServerSocket, addr string, port int32) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		log.Error("无法解析 Server 地址: " + err.Error())
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Error("无法监听 Server 地址: " + err.Error())
		return
	}

	socketChan := make(chan *TClientSocket)

	s.TCPListener = listener
	s.active = true
	s.activeConnections = 0

	go s.socketProducer(iSocket, socketChan)
	go s.socketConsumer(iSocket, socketChan)
}

func (s *TServerSocket) Active() bool {
	return s.active
}

func (s *TServerSocket) ActiveConnections() int {
	return s.activeConnections
}

func (s *TServerSocket) Connections() []*TClientSocket {
	return s.connections
}

func (c *TClientSocket) SocketHandle() uintptr {
	return c.socketHandle
}

func (c *TClientSocket) messageProducer(iSocket IClientSocket, conn *TClientSocket) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	var dataBuffer bytes.Buffer
	reading := false

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				vcl.ThreadSync(func() {
					iSocket.ClientSocketError(conn, err)
				})
			} else {
				conn.Close()
				vcl.ThreadSync(func() {
					iSocket.ClientSocketDisconnect(conn, err)
				})
			}
			break
		}

		for i := 0; i < n; i++ {
			if buffer[i] == '%' {
				reading = true
				dataBuffer.Reset()
				dataBuffer.WriteByte(buffer[i])
				continue
			}

			if buffer[i] == '$' {
				reading = false
				dataBuffer.WriteByte(buffer[i])
				message := dataBuffer.String()
				vcl.ThreadSync(func() {
					iSocket.ClientSocketRead(conn, message)
				})
				dataBuffer.Reset()
				continue
			}

			if reading {
				dataBuffer.WriteByte(buffer[i])
			}
		}
	}
}

func (c *TClientSocket) Dial(iSocket IClientSocket, addr string, port int32) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		log.Error("无法解析 Client 地址: " + err.Error())
		return
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error("无法监听 Client 地址: " + err.Error())
		return
	}

	c.TCPConn = tcpConn
	c.socketHandle = uintptr(unsafe.Pointer(tcpConn))

	go vcl.ThreadSync(func() {
		iSocket.ClientSocketConnect(c)
	})

	go c.messageProducer(iSocket, c)
}

func (c *TClientSocket) Close() {
	if c.TCPConn != nil {
		c.TCPConn.Close()
	}
}

func (c *TClientSocket) Write(message []byte) {
	if c.TCPConn != nil {
		c.TCPConn.Write(message)
	}
}

func (u *TUdpSocket) messageProducer(iUSocket IUdpSocket) {
	buffer := make([]byte, 2048) // 最大2048个字节
	for {
		numberBytes, _, err := u.ReadFromUDP(buffer)
		if err != nil {
			log.Error("读取UDP数据出错: {}", err.Error())
			vcl.ThreadSync(func() {
				iUSocket.UdpSocketError(u, err)
			})
			return
		}

		message := string(buffer[:numberBytes])
		vcl.ThreadSync(func() {
			iUSocket.UdpSocketRead(u, message)
		})
	}
}

func (u *TUdpSocket) ListenUDP(iUSocket IUdpSocket, serverAddr string, serverPort int32) {

	// 初始化UDP组件
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", serverAddr, serverPort))
	if err != nil {
		log.Error("无法解析地址: {}", err.Error())
		return
	}

	udpConn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Error("无法监听UDP: {}", err.Error())
		return
	}

	u.UDPConn = udpConn

	// 接收UDP消息
	go u.messageProducer(iUSocket)
}
