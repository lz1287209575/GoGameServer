package NetInterface

import "net"

// IConnection 定义连接接口
type IConnection interface {
	// Start 启动连接
	Start()

	// Stop 停止连接
	Stop()

	// GetTcpConnection 获取当前连接原始的Socket对象
	GetTcpConnection() *net.TCPConn

	// GetConnectionId 获取当前连接ID
	GetConnectionId() uint32

	// RemoteAddress 获取远程客户端地址信息
	RemoteAddress() net.Addr
}

// HandlerFunction 定义一个统一处理连接业务的接口
type HandlerFunction func(*net.TCPConn, []byte, int) error
