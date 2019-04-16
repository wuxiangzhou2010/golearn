package main

import "fmt"

type MyInterface interface {
	Print()
}
type MyStruct struct{}
func (ms MyStruct) Print(){}

func main() {
	x:=1
	var y interface {} = x //eface
	var s MyStruct
	var t MyInterface =s //iface
	fmt.Println(y, t)
}

/*

可以看出中间过程编译器将根据我们的转换目标类型的 empty interface 还是 non-empty interface，
来对原数据类型进行转换（转换成 <*_type, unsafe.Pointer> 或者 <*itab, unsafe.Pointer>）。
$ go build -gcflags="-N -l" -o main main.go //  disable compiler optimizations and inlining,
$ go tool objdump -s "main\.main" main.exe
*/