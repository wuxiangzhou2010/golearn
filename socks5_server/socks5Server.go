package main

// 1. demo for socks5 proxy protocol

//baidu.com for example
// test command: http_proxy="socks5://127.0.0.1:8081" curl 123.125.115.110

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

func main() {

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
	n, err := client.Read(b[:]) //read client fist message --> version/identity selection message
	if err != nil {
		log.Println(err)
		return
	}
	// 1. client version selection message
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+

	fmt.Println("1. ==> client send version selection message ==> ", b[:10])
	if b[0] == 0x05 { //socks protocol version
		methodselresp := []byte{0x05, 0x00} //method selection message, version5, no autentication
		fmt.Println("2. <== server response version selection message <== ", methodselresp)
		client.Write(methodselresp)
		n, err = client.Read(b[:]) //read client request
		if n == 0 {
			return
		}
		var host, port string
		fmt.Println("3. ==> client send request==> ", b[:10])
		switch b[3] { // address type ipv4/domain name/ipv6
		case 0x01:
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03:
			host = string(b[5 : n-2])
		case 0x04:
			host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()

		}
		fmt.Println(n)
		port = strconv.Itoa(int(b[n-2])<<8 | int(b[n-1])) //port
		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		fmt.Printf("host: %s, port %s\n", host, port)
		if err != nil {
			log.Println(err)
			return
		}
		defer server.Close()
		// version| response | reserved| address type
		connectionresp := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		fmt.Println("4. <== server send connection status message <== ", connectionresp)
		client.Write(connectionresp)
		go io.Copy(server, client)
		io.Copy(client, server)
	}

}
