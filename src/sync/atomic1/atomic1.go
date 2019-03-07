package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("counting to 50000 ...")

	test1()
	test2()
}

func test1() {
	var wg sync.WaitGroup
	count := 0

	t := time.Now()
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			count++
			wg.Done()
		}(&wg, i)

	}
	wg.Wait()
	fmt.Println(time.Now().Sub(t))
	fmt.Println("count====> ", count)
	fmt.Println("exit")
}

func test2() {
	var wg sync.WaitGroup
	// wg2 := wg
	// wg = wg2
	// nocopy
	//  $ go vet atomic1.go
	// # command-line-arguments
	// .\atomic1.go:36:9: assignment copies lock value to wg2: sync.WaitGroup contains sync.noCopy
	// .\atomic1.go:37:7: assignment copies lock value to wg: sync.WaitGroup contains sync.noCopy

	count := int32(0)

	t := time.Now()
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			atomic.AddInt32(&count, 1)
			wg.Done()
		}(&wg, i)

	}
	wg.Wait()
	fmt.Println(time.Now().Sub(t))
	fmt.Println("count====> ", count)
	fmt.Println("exit")
}
