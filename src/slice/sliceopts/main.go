package main

import "fmt"

func printslice(s []int) {
	fmt.Println("len, cap: ", s, "\n", len(s), cap(s))
}
func main() {
	//1. creating slice
	//var s []int
	fmt.Println("1. creating slice")
	s := make([]int, 0, 8) // 每次reslice , cap 会变成原来的两倍
	for i := 0; i < 20; i++ {
		s = append(s, 2*i+1) // reslice 会导致slice地址变化
		fmt.Println("len and cap ", len(s), cap(s))
	}
	printslice(s)
	//2. copy slice
	fmt.Println("2. copying slice")
	s2 := make([]int, 30)
	copy(s2, s)
	printslice(s2)

	//3. delete slice
	fmt.Println("3. delete slic 3")
	s3 := append(s[:3], s[4:]...)

	printslice(s3)
	// 4. poping front
	fmt.Println("4. poping front")
	front := s3[0]
	fmt.Println("front ", front)
	s3 = s[1:]

	printslice(s3)
	// 5. poping from back
	fmt.Println("5. poping from back")
	tail := s3[len(s3)-1]
	fmt.Println("tail ", tail)
	printslice(s3[:len(s3)-1])

}
