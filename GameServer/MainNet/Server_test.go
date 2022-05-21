package MainNet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest(t *testing.T) {
	fmt.Println("Client Test ... Start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("Client Start Error, Exit", err)
		return
	}

	for i := 0; i < 3; i++ {
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

	err = conn.Close()
	if err != nil {
		return
	}
	t.Logf("Successful")
}

func TestServer(t *testing.T) {
	s := NewServer("GoGameServer", "0.0.0.0", 7777)

	go ClientTest(t)

	s.MainLoop()
}
