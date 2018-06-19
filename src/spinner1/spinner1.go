package main

import (
	"fmt"
	"time"
)

func spinner(d time.Duration) {
	for {
		for _, s := range `\|/` {
			fmt.Printf("\r%c", s)
			time.Sleep(d)
		}
	}
}

func main() {
	go spinner(time.Millisecond * 199)
	const f = 40
	fmt.Print("\r", fib(f))
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
