package main

import (
	"fmt"
	"time"
)


func main(){
	myChannel := make(chan int, 1)
	num := 6
	go func(){
		sender := myChannel
		sender <- num
		fmt.Println("sent!")
		
	}()

	go func(){
		fmt.Println("received!",<-myChannel)
	}()

	time.Sleep(time.Second)
}
