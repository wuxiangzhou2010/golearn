package main

// 1. interface hold type and value, and it can be nil
import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}
func main() {
	var i I  // nil interface, both value and type is nil
	var t *T // nil *T
	i = t    // i's type is not nil, but the value is nil
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

}
