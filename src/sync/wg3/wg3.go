package main

import (
	"fmt"
	"sync"

	"./game"
)

var wiatGroup sync.WaitGroup

func main() {
	wiatGroup.Add(1)
	serverAddr := "baidu.com:80"
	go game.ConnSocket(serverAddr, &wiatGroup)
	wiatGroup.Wait()
	fmt.Println("main Done")
}
