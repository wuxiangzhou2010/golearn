package main

import "fmt"

func main() {
	num := 5
	if num += 4; 10 > num {
		num += 3
		fmt.Println(num)
	} else if 10 < num {
		num -= 2
		fmt.Println(num)
	}
	fmt.Println(num)
}
