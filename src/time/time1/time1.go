package main

// 1. select from  channel and default
// 2. use time.Sleep to delay

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "Process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}
