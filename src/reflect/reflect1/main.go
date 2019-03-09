package main

// 1. struct if value type
// 2. xxx.(type) can be used in a type switch
// 3. reflect.TypeOf can get the type of a variable

import (
	"fmt"
	"reflect"
)

type ntt struct {
	name string
	data interface{}
}

type xx struct {
	name  string
	value int
}

func main() {

	n1 := &ntt{
		name: "test",
		data: "string",
	}

	x := xx{
		name:  "xx",
		value: 33,
	}

	n2 := &ntt{
		name: "test",
		data: x,
	}
	n3 := &ntt{
		name: "test",
		data: &x,
	}

	tt(n1)
	tt(n2)
	tt(n3)

}

func tt(ntt *ntt) {
	switch ntt.data.(type) {

	case nil:
		fmt.Println("case nil: ", ntt.data)
	case xx:
		var yy xx
		yy = ntt.data.(xx)
		fmt.Println("case xxx: ", yy)
	default:
		fmt.Println("case default:", reflect.TypeOf(ntt.data))
	}
}
