package main

import "fmt"

//1 . assign a concrete type to a interface type, the concrete should implement the interface

// I is an interface
type I interface {
	M()
}

// T is a struct
type T struct {
	S string
}

// M is implemented by T,so T implement interface I
func (t T) M() {
	fmt.Println(t.S)

}

func main() {
	var i I = T{"hello"} // instance of T can be assigned to interface i directly
	i.M()

}
