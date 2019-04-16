package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i //panic: send on closed channel
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("closed")
				return
			}
			fmt.Println("a: ", a)
		}

	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)

}

/*
reference: https://zhuanlan.zhihu.com/p/35058068 题目5
*/
