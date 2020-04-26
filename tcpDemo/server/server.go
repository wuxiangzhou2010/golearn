package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("connection read error:", err)
		return
	}
	fmt.Println("received:", string(buf[:n]))
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Listen :20000 failed, err:", err)
		return
	}
	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println("connection failed, err:", err)
			continue
		}

		go process(conn)
	}

}
