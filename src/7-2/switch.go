package main

// 1. change seed for random number generator
// 2. type assertion
// 3. rand.Seed(int64),  rand.Intn(int) int
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)} //interface{} array can hold any type of variable
	rand.Seed(time.Now().UnixNano())                       // change seed, int64
	// switch v := ia[rand.Intn(4):]; interface{}(v).(type) { // case A
	switch v := ia[rand.Intn(4)]; interface{}(v).(type) { // case B C D  default
	case []interface{}:
		fmt.Println("case A")
	case byte:
		fmt.Println("case B")
	case string:
		fmt.Println("case C")
	case uint:
		fmt.Println("case D")
	// case int32:
	// 	fmt.Println("case E")
	default:
		fmt.Println("Unknown!")
	}
}
