package main

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		msg.Name // this is wrong
	}
}

func main() {
	s := &student{"student"}
}

/*
reference: https://zhuanlan.zhihu.com/p/35058068 第二题

*/
