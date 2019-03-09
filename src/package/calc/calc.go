package main

//1. package stconv Atoi
//2. switch xxx {}
//3. os.Args
//4. testing

// refer: https://studygolang.com/articles/5831
import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguements]...")
	fmt.Println("\nThe commands are\n\tadd\taddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <interger1> <interger2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <interger1> <interger2>")
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)

	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <interger>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <interger>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()

	}
}