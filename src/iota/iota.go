package main

import "fmt"

func main() {
	const (
		num1 int = iota + 1
		num2
		num3
	)
	fmt.Println(num3)
}
