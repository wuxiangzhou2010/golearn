package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json: "name"` //  小写为私有成员，json解码无法访问, people:{}
}

func main() {
	js := `{
		"name": "11"
		}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("People: ", p)

}

/*
reference: https://zhuanlan.zhihu.com/p/35058068 题目3
*/
