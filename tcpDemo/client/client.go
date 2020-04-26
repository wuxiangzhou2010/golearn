package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("connection error:", err)
		return
	}

	_, err = conn.Write([]byte("how are you?"))
	if err != nil {
		fmt.Println("send error:", err)
	}

}
