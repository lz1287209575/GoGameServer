package MainNet

import (
	"fmt"
	"net"

	"github.com/louxiaoche/go-game-server/GameServer/NetInterface"
)

type Connection struct {
	// 当前连接的Tcp套接字
	Conn *net.TCPConn

	// 当前连接的ID 也可以称作SessionID, ID全局唯一
	ConnectionId uint32

	// 当前连接的关闭状态
	IsClosed bool

	// 该连接的处理方法API
	HandlerAPI NetInterface.HandlerFunction

	// 告知该连接已经退出/停止的Channel
	ExitBufferChannel chan bool
}

// NewConnection 创建新连接对象的函数
func NewConnection(conn *net.TCPConn, connectionId uint32, callbackApi NetInterface.HandlerFunction) *Connection {
	c := &Connection{
		Conn:              conn,
		ConnectionId:      connectionId,
		IsClosed:          false,
		HandlerAPI:        callbackApi,
		ExitBufferChannel: make(chan bool, 1),
	}
	return c
}

// StartReader 处理conn读取数据的Goroutine
func (conn *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running ... ")
	defer fmt.Println(conn.RemoteAddress().String(), " conn reader exit!")
	defer conn.Stop()
	for {
		// 读取我们最大的数据到buffer当中
		buf := make([]byte, 512)
		count, err := conn.Conn.Read(buf)
		if err != nil {
			fmt.Println("Receive Buffer Error", err)
			conn.ExitBufferChannel <- true
			continue
		}

		// 调用当前连接业务
		if err := conn.HandlerAPI(conn.Conn, buf, count); err != nil {
			fmt.Println("ConnectionId: ", conn.ConnectionId, " handle is error")
			conn.ExitBufferChannel <- true
			return
		}
	}
}

// Start 启动连接，当前连接开始工作
func (conn *Connection) Start() {
	go conn.StartReader()
	for {
		select {
		case <-conn.ExitBufferChannel:
			// 得到退出消息，不再阻塞
			return
		}
	}
}

// Stop 停止链接，结束当前状态
func (conn *Connection) Stop() {
	// 1. 当前连接已经关闭
	if conn.IsClosed {
		return
	}

	conn.IsClosed = true

	// TODO: Connection Stop() 如果用户注册了该连接的关闭回调业务，那么此刻应该是显示调用
	// 关闭Socket连接
	conn.Conn.Close()
	// 通知从缓冲队列读取数据的业务，该连接已经关闭
	conn.ExitBufferChannel <- true
	close(conn.ExitBufferChannel)
}

// GetTcpConnection 从当前连接获取原始的Socket对象 TCPConn
func (conn *Connection) GetTcpConnection() *net.TCPConn {
	return conn.Conn
}

// GetConnectionId 获取当前连接ID
func (conn *Connection) GetConnectionId() uint32 {
	return conn.ConnectionId
}

// RemoteAddress 获取远程客户端地址信息
func (conn *Connection) RemoteAddress() net.Addr {
	return conn.Conn.RemoteAddr()
}
