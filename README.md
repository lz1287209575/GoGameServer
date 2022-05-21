# GoGameServer

## v0.1版本 基础的Server框架

1. 启动服务器
   1. 创建Address
   2. 创建Listener
   3. 处理基础客户端业务
2. 停止服务器
   1. 资源回收和状态回执
3. 运行服务器
   1. 调用Start函数后阻塞处理
4. 初始化Server


## v0.2版本 链接封装和业务绑定

1. 启动链接
2. 停止链接
3. 获取当前链接的conn对象
4. 得到链接ID
5. 得到客户端链接的地址和端口
6. 发送数据
7. 链接所绑定的处理业务的函数模型

--- 

- Socket Tcp套接字
- 链接的ID ConnID
- 当前链接状态 isClosed
- 与当前链接绑定的处理业务方法 HandleAPI
- 等待链接被动退出 channel ExitChannel