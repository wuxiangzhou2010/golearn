package main

import (
	"fmt"
	"sync"
)

// func main() {

// 	m := make(map[int]int)

// 	var wg sync.WaitGroup
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(index int) {

// 			m[i] = i
// 			wg.Done()
// 		}(i)
// 	}

// 	wg.Wait()

// }

func main() {

	m := sync.Map{}

	var wg sync.WaitGroup
	for i := 0; i < 20000000; i++ {
		wg.Add(1)
		go func(index int) {

			m.Store(index, index) // key value

			if v, ok := m.Load(index); ok {
				fmt.Println(v)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

}
