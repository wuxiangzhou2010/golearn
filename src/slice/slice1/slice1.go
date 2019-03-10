package main

import "fmt"

type tx struct {
	i int
	s string
}

type txs []*tx

func modify(s txs) {
	for _, v := range s {

		v.i = 222
	}
}

func main() {

	t1 := &tx{
		1,
		"first",
	}
	t2 := &tx{
		2,
		"second",
	}

	s := make(txs, 0)
	s = append(s, t1)
	s = append(s, t2)

	modify(s)

	for i, v := range s {
		fmt.Println("index , v ", i, v)
	}
}
