package main

//1. use channel to get fibonacci array
//2. Do only once for every fib number

import "fmt"

func main() {
	ch := make(chan int)
	go fibonacci(10, ch)

	for i := range ch {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
