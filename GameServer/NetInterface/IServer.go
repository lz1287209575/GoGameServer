package NetInterface

type IServer interface {
	// Start 启动服务器
	Start()

	// Stop 停止服务器
	Stop()

	// MainLoop 服务器主循环
	MainLoop()
}
