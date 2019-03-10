package main

// 1. a type can implement an interface by value or by pointer, they can not be mixed

import (
	"fmt"
	"math"
)

// I is an interface, has M() method
type I interface {
	M()
}

// T is a struct has string
type T struct {
	S string
}

// M method is implemented by *T
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// M method is implemented by F
func (f F) M() {
	fmt.Println(f)
}
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	i = &T{"hello"} //ok
	// i = T{"hello"}  // error if use i=T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}
