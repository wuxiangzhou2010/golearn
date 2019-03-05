package main

import "fmt"

// slice 相当于一个包含{指向数组指针， slice 占用的长度， 数组容量}的结构体，
// 对slice 的copy 操作不会改变原slice大小和指向关系， 但是内容可能改变
func main() {
	// a := []string{"a", "b", "c", "d"}
	a := make([]int, 10, 30)
	a = a[:0]
	app(a, 1)
	app(a, 2)

	fmt.Println(a, len(a), cap(a))
	// fmt.Println(a[2:8])
	fmt.Println(a[:20])

}

func app(a []int, i int) {
	a = append(a, i)
}
