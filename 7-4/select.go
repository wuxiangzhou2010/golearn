package main

import "fmt"

func main() {
	ch := make(chan string, 3)// 3 string at most 
	ch<- "test1"
	ch<- "test2"
	ch<- "test3"
	for i := 0; i < 4; i++ {
		select {
		case e, ok := <-ch:
			if !ok {
				fmt.Println("End")
				return
			}
			fmt.Println(e)
		default:
			fmt.Println("NO data")
		}
	}
}
