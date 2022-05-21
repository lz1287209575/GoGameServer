package MainNet

// import 注意顺序问题，标准库 - 项目 - 外部
import (
	"fmt"
	"net"
	"time"

	"github.com/louxiaoche/go-game-server/GameServer/NetInterface"
)

// IServer 接口实现

type Server struct {
	// 服务器名称
	Name string

	// IPv4 or other
	IPVersion string

	// IP地址
	IPAddress string

	// 端口
	Port int
}

func CallbackApi(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显业务
	fmt.Println("[Connection Handler] CallbackApi ...... ")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("Write Back Buffer Error ", err)
		return err
	}
	return nil
}

func (s *Server) Start() {
	// 实现服务器启动一共有三个步骤
	// 1. 获取Tcp地址
	// 2. 监听该地址
	// 3. 启动server 网络连接业务
	fmt.Printf("[START] Server Listener at IP: %s, Port: %d is Starting", s.IPAddress, s.Port)
	// 创建一个go协程
	go func() {
		// 获取一个TCP的地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IPAddress, s.Port))

		if err != nil {
			fmt.Println("Resolve Tcp Addr Error: ", err)
			return
		}

		// 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)

		if err != nil {
			fmt.Println("Listen ", s.IPVersion, "Error: ", err)
			return
		}

		// 已经监听成功
		fmt.Println("Start Server: ", s.Name, "Successful, Now listening ......")

		// TODO: Server.go里面应该有自动生成ID的接口

		var connectionId uint32
		connectionId = 0

		// 启动Server网络连接业务
		for {
			// 阻塞等待客户端建立连接
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept Error ", err)
				continue
			}

			// TODO: Server.Start() 设置服务器最大连接控制，超过阈值则拒绝客户端连接
			// TODO: Server.Start() 处理该连接的请求 业务方法 handler应该和connection是绑定的

			// 先暂时做一个最大512字节的回显服务
			connection := NewConnection(conn, connectionId, CallbackApi)
			connectionId++
			go connection.Start()
		}

	}()
}

func (s *Server) Stop() {
	// 先仅仅实现打印，这里需要其他的功能协助
	fmt.Printf("[Stop] Server: %s is Shutdown", s.Name)
}

func (s *Server) MainLoop() {
	s.Start()
	// 先仅仅实现打印，这里需要其他的功能协助
	for {
		time.Sleep(10 * time.Second)
	}
}

// 创建新的服务器句柄

func NewServer(name string, ipAddr string, port int) NetInterface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IPAddress: ipAddr,
		Port:      port,
	}
	return s
}
