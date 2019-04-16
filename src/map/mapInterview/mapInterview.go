package main

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu" // cannot assign to struct field m["people"].name in mapgo
}

/*
reference:
https://zhuanlan.zhihu.com/p/35058068 issue 9
*/
