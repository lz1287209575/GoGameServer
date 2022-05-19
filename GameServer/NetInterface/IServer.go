package NetInterface

type IServer interface {
	// 启动服务器
	Start()

	// 停止服务器
	Stop()

	// 服务器主循环
	MainLoop()
}
