package main

import "fmt"

type People interface {
	Name () string
}

type Student struct {name string}

func (stu * Student) Name() string{
	return stu.name
}
func getPeople() People {
	var stu * Student
	return stu
}

func main() {
	p := getPeople()
	if p== nil { // interface is something like a struct , the interface value is nil but the interface type is not nil
		fmt.Println("AAA")
	} else {
		fmt.Println("BBB")
	}
}