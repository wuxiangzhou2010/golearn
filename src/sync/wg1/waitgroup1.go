package main

import (
	"fmt"
	"sync"
)

// goroutine 并行访问一个map, 最终打印的顺序是随机的。

func main() {

	// WaitGroup
	a := make(map[int]string)
	a[1] = "aaa"
	a[2] = "bbb"
	var wg sync.WaitGroup

	for k, v := range a {
		wg.Add(1)
		go func(k int, v string, wg *sync.WaitGroup) {
			fmt.Println("k ", k, "v ", v)
			defer wg.Done()
		}(k, v, &wg)
	}
	wg.Wait()
}
