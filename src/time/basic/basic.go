package main

//1. time.Hour(), time.Now()
import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	switch {

	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
