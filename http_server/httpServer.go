package main

// 1. basic http proxy

//  test command: http_proxy="socks5://127.0.0.1:8081" wget sina.com:80
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	var method, host, address string
	fmt.Println("1. ==> firt message from the client==>")
	fmt.Println("===============================\n", string(b[:128]), "\n===============================")
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	fmt.Printf("method = %s,host = %s\n", method, host)
	hostPortURL, err := url.Parse(host)

	if err != nil {
		log.Println(err)
		return
	}
	// parse server port and address
	if hostPortURL.Opaque == "443" {
		address = hostPortURL.Scheme + ":443"
	} else {
		if strings.Index(hostPortURL.Host, ":") == -1 {
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}

	// if CONNECT method
	if method == "CONNECT" {
		connectrespmsg := "HTTP/1.1 200 Connection established\r\n\r\n"
		fmt.Fprint(client, connectrespmsg)
		fmt.Println("<== server respond CONNECT method with  ", connectrespmsg)
	} else {
		//command  like GET ...
		server.Write(b[:n])
	}

	go io.Copy(server, client)
	io.Copy(client, server)
}
