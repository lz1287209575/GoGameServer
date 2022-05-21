package main

// 引用当前项目的目录，带上gomod里面的module路径就可以了，另外一般包名都是小驼峰，文件名是下划线，尽量使用单个单词描述
import "github.com/louxiaoche/go-game-server/GameServer/MainNet"

func main() {
	server := MainNet.NewServer("GoGameServer", "127.0.0.1", 7777)
	server.MainLoop()
}
