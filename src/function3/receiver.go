package main

//  1. pointer receiver is copied
import "fmt"

type test struct {
	name string
}

func (t test) TestValue() {
	fmt.Printf("in func (t test) TestValue() t  %p\n", &t)

}
func (t *test) TestPointer() {
	fmt.Printf("in func (t *test) TestPointer t %p\n", t)

}

func main() {
	t := test{}
	fmt.Printf("pointer of t %p\n", &t)
	m := test.TestValue
	m(t)
	m1 := (*test).TestPointer
	m1(&t)

}
