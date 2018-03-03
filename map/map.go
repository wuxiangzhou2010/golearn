package main

import "fmt"

func main(){
	mm2 := map[string]int{
		"golang": 42,
		"java":1,
		"python": 8,
	}
	fmt.Printf("%d, %d, %d \n", mm2["golang"],mm2["python"],mm2["java"])
}