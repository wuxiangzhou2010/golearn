package main

//==============================
// 1. generate in range [0,n)
// 2. will panic when rand.Intn(0)

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	fmt.Println(rand.Intn(0))
// 	// fmt.Println(rand.Intn(1))
// }

//==============================
// 1. get a slice from random number gnerator

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n) from the default Source.
	i := rand.Perm(10)
	fmt.Printf("%T, %v ", i, i)
	// for _, k := range rand.Perm(10) {
	// 	fmt.Println(k)
	// }
}

//==============================
