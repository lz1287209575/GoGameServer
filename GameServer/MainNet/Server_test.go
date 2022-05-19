package MainNet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("Client Test ... Start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("Client Start Error, Exit", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello GoGameServer"))
		if err != nil {
			fmt.Println("Write Error err", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read Buffer Error ", err)
			return
		}

		fmt.Printf(" Server Callback: %s, cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}

}

func TestServer(t *testing.T) {
	s := NewServer("GoGameServer", "0.0.0.0", 7777)

	go ClientTest()

	s.MainLoop()
}
