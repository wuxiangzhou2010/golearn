package main

import "fmt"

// slice 相当于一个包含{指向数组指针， slice 占用的长度， 数组容量}的结构体，
// 对slice 的copy 操作不会改变原slice大小和指向关系， 但是内容可能改变
func main() {
	// a := []string{"a", "b", "c", "d"}
	a := make([]int,10, 60)
	a = a[:0]
	app(a, 1) // 1被添加到了数组的第一个元素
	fmt.Println(a[:20], len(a), cap(a))
	app(a, 2) // 数组第一个元素被覆盖， 变为了2

	fmt.Println(a, len(a), cap(a))
	// fmt.Println(a[2:8])
	fmt.Println(a[:20])


}
/*
result:
[] 0 30
[2 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

*/



func app(a []int, i int) {
	a = append(a, i)
}
