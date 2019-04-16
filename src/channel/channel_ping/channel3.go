package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		messages <- "string"
	}()
	msg := <-messages
	fmt.Println(msg)
}
