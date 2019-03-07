package game

import (
	"fmt"
	"net"
	"sync"
)

var gameWait *sync.WaitGroup

func ConnSocket(serverAddr string, wait *sync.WaitGroup) {
	var err error
	Conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connected:", err)
		wait.Done()
		return
	}
	gameWait = wait
	fmt.Println("connected Ok:", Conn.RemoteAddr())
	go readMessage()
}

func readMessage() {
	fmt.Println("readMessage")
	for {
		if dosomething() {
			gameWait.Done()
			return
		}
	}
}

func dosomething() bool { return true }
