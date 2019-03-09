package main

// 1. use map reduce to get the sum of a slice
// 2. channel is thread safe
// https://golang.org/ref/spec#Channel_types

import "fmt"

func sum(s []int, c chan int) {
	var tmp int
	for _, v := range s {
		tmp += v
	}
	c <- tmp
}

func main() {
	s := []int{1, 2, 334, 45, 545, 6, 6, 3434, 23, 1}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println("sum of slice s: ", x+y)
}
