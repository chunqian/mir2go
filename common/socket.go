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
	IPaddr       int
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

	sockChan := make(chan *TClientSocket)

	msgProducer := func(iSocket IServerSocket, conn *TClientSocket) {
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
					iSocket.ServerSocketClientError(conn, err)
				} else {
					conn.Close()
					iSocket.ServerSocketClientDisconnect(conn, err)
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
					// log.Info("Message Received: {}", dataBuffer.String())
					message := dataBuffer.String()
					iSocket.ServerSocketClientRead(conn, message)
					dataBuffer.Reset()
					continue
				}

				if reading {
					dataBuffer.WriteByte(buffer[i])
				}
			}
		}
	}

	sockProducer := func(ch chan *TClientSocket) {
		for {
			conn, err := listener.Accept()
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
			iSocket.ServerSocketClientConnect(clientSocket)

			ch <- clientSocket
		}
	}

	sockConsumer := func(ch chan *TClientSocket) {
		for {
			select {
			case sock := <-ch:
				go msgProducer(iSocket, sock)
				// default:
				// 	log.Error("Could not read from sockChan")
			}
		}
	}

	s.TCPListener = listener
	s.active = true
	s.activeConnections = 0

	go sockProducer(sockChan)
	go sockConsumer(sockChan)
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

	msgProducer := func(iSocket IClientSocket, conn *TClientSocket) {
		defer conn.Close()

		buffer := make([]byte, 1024)
		var dataBuffer bytes.Buffer
		reading := false

		for {
			n, err := conn.Read(buffer)
			if err != nil {
				if err != io.EOF {
					iSocket.ClientSocketError(conn, err)
				} else {
					conn.Close()
					iSocket.ClientSocketDisconnect(conn, err)
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
					// log.Info("Message Received: {}", dataBuffer.String())
					message := dataBuffer.String()
					iSocket.ClientSocketRead(conn, message)
					dataBuffer.Reset()
					continue
				}

				if reading {
					dataBuffer.WriteByte(buffer[i])
				}
			}
		}
	}

	c.TCPConn = tcpConn
	c.socketHandle = uintptr(unsafe.Pointer(tcpConn))

	iSocket.ClientSocketConnect(c)

	go msgProducer(iSocket, c)
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
