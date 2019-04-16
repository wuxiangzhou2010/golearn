package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string { // 栈空间溢出
	return fmt.Sprintf("print :%v ", p) // pointerInterview.go:10:9: Sprintf format %v with arg p causes recursive String method call
}
func main() {
	p := &People{}
	p.String()
}

/*

reference: https://zhuanlan.zhihu.com/p/35058068 题目4



runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow
*/
