package main

// 1. struct is a concrete type
// 2. pointers to this struct can see the change

import "fmt"

type test struct {
	i int
	s string
}

func main() {

	var t1 []*test

	var t = test{
		i: 32,
		s: "test",
	}
	t1 = append(t1, &t)
	fmt.Println(t1[0].i)
	t.i = 0

	fmt.Println(t1[0].i)

}
