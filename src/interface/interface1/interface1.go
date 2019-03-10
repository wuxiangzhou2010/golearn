package main

import "fmt"

type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name  string
	Age   uint8
	Where string
}

//using receiver  implement  he interface
func (c *Cat) Grow() {
	fmt.Println("Grow")
}

//using receiver implement  he interface
func (c *Cat) Move(s string) string {
	fmt.Println("Move")
	return s
}

func main() {
	cat := Cat{"little c", 2, "In the house"}

	animal, ok := interface{}(&cat).(Animal)

	fmt.Printf("%v, %v. \n", animal, ok)
	animal.Grow()
	animal.Move("Beijing")

}
