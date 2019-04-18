package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan *string)
	go func() {
		line1 := "test is here"
		fmt.Println("in gouroutine", &line1)
		c <- &line1
	}()
	time.Sleep(5 * time.Second)
	v := <-c
	fmt.Println(v, *v)
}
