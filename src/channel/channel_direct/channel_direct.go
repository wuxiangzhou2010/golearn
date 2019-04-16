package main

import "fmt"

func ping(c chan string, s string) {
	c <- s
}

func pong(c <-chan string, pong chan<- string) {
	msg := <-c
	pong <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "hello")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
