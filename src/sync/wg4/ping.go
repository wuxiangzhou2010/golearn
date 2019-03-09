package main

// 1. multiple go routine read from the same channel
import (
	"fmt"
	"strconv"
	"sync"
)

func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- strconv.Itoa(i) + "ping"
	}
	close(c)
}

func main() {
	c := make(chan string, 3)

	go pinger(c)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for msg := range c {
			fmt.Println("r1 ", msg)
		}
		wg.Done()
	}()
	go func() {
		for msg := range c {
			fmt.Println("r2 ", msg)
		}
		wg.Done()
	}()
	wg.Wait()

}
