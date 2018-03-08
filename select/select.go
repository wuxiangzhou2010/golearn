package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := make(chan string)
	s2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		s1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		s2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-s1:
			fmt.Println(msg)
		case msg := <-s2:
			fmt.Println(msg)
		}

	}

}
