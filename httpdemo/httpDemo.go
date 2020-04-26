package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "aidu.com:80")
	defer conn.Close()
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}

	n, err := conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	if err != nil {
		fmt.Println("write to tcp error:", err)
		return
	}

	var buf [1024]byte

	for {
		n, err = conn.Read(buf[:])
		if err == io.EOF {
			fmt.Printf("%s", buf[:n])
			return
		}
		if err != nil {
			fmt.Println("read error:", err)
			return
		}
		fmt.Printf("%s", buf[:n])

	}

}