package main

type Param map[string]interface{}

type Show struct {
	Param // map必须make 否则为nil
}

func main() {
	s := new(Show)
	s.Param["RMB"] = 10000 // panic: assignment to entry in nil map
}

/*
reference:
https://zhuanlan.zhihu.com/p/35058068  issue1
*/
