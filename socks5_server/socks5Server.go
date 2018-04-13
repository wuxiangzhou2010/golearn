package main

import (
	"io"
	"log"
	"net"
	"fmt"
	"strconv"
)


func main() {
	//log.SetFlags(log. | log.Lshortfile)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()

		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}

	defer client.Close()

	var b [4096]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	if b[0] == 0x05 {
		client.Write([]byte{0x05, 0x00})
		n, err = client.Read(b[:])
		if n == 0 {
			return
		}
		var host, port string

		switch b[3] {
		case 0x01:
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03:
			host = string(b[5: n - 2])
		case 0x04:
			host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()

		}
		fmt.Println(n)
		port = strconv.Itoa(int(b[n-2])<<8 | int(b[n-1]))
		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		fmt.Printf("host: %s, port %s\n", host, port)
		if err != nil {
			log.Println(err)
			return
		}
		defer server.Close()
		client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,0x00,0x00, 0x00})
		go io.Copy(server, client)
		io.Copy(client,server)
	}

}