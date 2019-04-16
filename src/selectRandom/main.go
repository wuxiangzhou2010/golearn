package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	string_chan <- "let's go"

	int_chan <- 1
	select {
	default:
		fmt.Println("default")
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		fmt.Println(value)
	}

}
