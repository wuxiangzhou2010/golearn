package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	switch v := ia[rand.Intn(4):]; interface{}(v).(type) {
	case []interface{}:
		fmt.Println("case A")
	case byte:
		fmt.Println("case B")
	default:
		fmt.Println("Unknow!")
	}

}
