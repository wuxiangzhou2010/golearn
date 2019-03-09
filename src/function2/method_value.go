package main

//1. the method of a concrete type or a pointer type can be assigned or even be passed as an argument
import "fmt"

type test struct {
	name string
}

func (t test) TestValue() {
	fmt.Println("call from func (t test) TestValue()")

}
func (t *test) TestPointer() {
	fmt.Println("call from func (t *test) TestPointer()")
}

func main() {
	t := test{}
	m := test.TestValue // function of

	m(t)
	m1 := (*test).TestPointer
	m1(&t)

	fmt.Printf("type of m %T\ntypep of m1 %T", m, m1)

}
