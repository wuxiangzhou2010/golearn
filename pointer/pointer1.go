package main

import "fmt"

type MyInt struct {
	n int
}

func (i *MyInt) Increase() {
	i.n++
	fmt.Println(i.n)
}

func (i *MyInt) Decrease() {
	i.n--
	fmt.Println(i.n)
}

func main() {
	t := MyInt{1}
	t.Increase()
	t.Increase()
	t.Increase()
	t.Decrease()
	t.Decrease()
	t.Decrease()

}
