package main

// 1. only interface type can use type asssertion

import (
	"fmt"
)

func main() {

	a := interface{}(20)
	b := "1212323"
	c := 1212.222
	//1. ok
	i := a.(int)

	//1. panic, a's underlying type is not a string
	// tmp := a.(string)
	// fmt.Println("tmp ", tmp)

	//2. prevent panic
	_, ok := a.(string)
	if ok == false {
		fmt.Println("a is not  a string")
	}

	//3. non-interface type int on left
	// d := 30
	// j := d.(int)

	fmt.Printf("type of \na\tb\tc\ti\n%T\t%T\t%T\t%T", a, b, c, i)
}
