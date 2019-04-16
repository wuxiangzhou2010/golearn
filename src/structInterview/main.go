package main

type Param map[string]interface{}

type Show struct {
	Param  // map必须make 否则为nil
}

func main() {
	s := new(Show)
	s.Param["RMB"] = 10000
}