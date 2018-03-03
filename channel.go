package main

import "fmt"

func main(){
	c := make(chan string, 6)
	go func(){
		c <- "hello"
	}()
	v := <-c
	fmt.Println(v)
}