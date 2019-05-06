package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func parseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for index, stu := range stus { // index 和value的地址是不变的。
		fmt.Printf("index %p value %p\n", &index, &stu)
		m[stu.Name] = &stu
	}


}

func main()  {
	parseStudent()
}
