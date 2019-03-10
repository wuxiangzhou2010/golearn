package main

//1.  defer must be a function, either a concrete function call or an annoymous function call.
//2.  a function can have multiple defer statement, and they will be execute in Last In First Out manner.
//3.  defer will be execute after the surround function returns, but before the return value is assigned to the variable.

import "fmt"

func main() {

	fmt.Println("in main ", *increase1())
}

func increase1() *int {
	i := 100
	defer increase2(&i) // 103
	defer increase2(&i) // 102
	defer increase2(&i) // 101
	defer func() { fmt.Println("third") }()
	defer func() { fmt.Println("second") }()
	defer func() { fmt.Println("first") }()

	fmt.Println("in increase1", i)
	return &i

}
func increase2(i *int) {
	*i = *i + 1
	fmt.Println("in increase2", *i)
}
