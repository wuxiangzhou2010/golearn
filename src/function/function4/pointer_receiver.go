package main

// 1.execute  function on a varibale can be delayed
import "fmt"

type test struct {
	name string
}

func (t test) TestValue() {
	fmt.Printf("%s\n", t.name)

}

func (t *test) TestPointer() {
	fmt.Printf("%s\n", t.name)

}

func main() {
	t := test{name: "wang"}
	m := t.TestValue // function on t(copy), but not execute
	t.name = "Li"
	m1 := (*test).TestPointer //function
	m1(&t)                    // execute funcion on &t
	m()                       // function on t executed
}
