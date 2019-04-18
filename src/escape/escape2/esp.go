//example1
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	s := "hello"
// 	fmt.Println(s)
// }

// Example2:

// package main

// type S struct{}

// func main() {
// 	var x S
// 	_ = *ref(x)
// }

// func ref(z S) *S {
// 	return &z
// }

// example3
// package main

// type S struct {
// 	M *int
// }

// func main() {
// 	var i int
// 	refStruct(i)
// }

// func refStruct(y int) (z S) {
// 	z.M = &y
// 	return z
// }

// Example4:

// package main

// type S struct {
// 	M *int
// }

// func main() {
// 	var i int
// 	refStruct(&i)
// }

// func refStruct(y *int) (z S) {
// 	z.M = y
// 	return z
// }

// Example5:

package main

type S struct {
	M *int
}

func main() {
	var x S
	var i int
	ref(&i, &x)
}

func ref(y *int, z *S) {
	z.M = y
}
