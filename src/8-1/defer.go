package main

import "fmt"

func main() {
	for i := uint(0); i < 10; i++ {
		fmt.Println(fb(i))
	}
}

func fb(n uint) uint {
	if n < 2 {
		return 1
	}

	return n + fb(n-1)
}
