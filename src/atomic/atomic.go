package main

/*
demonstrate  the atomic.CompareAndSwapInt32 fucntion
*/

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func c32To42(total *int32, id int) {
	if atomic.CompareAndSwapInt32(total, 32, 42) {
		fmt.Printf("32 to 42: id %d, work\n", id)
	} else {
		fmt.Printf("32 to 42:id %d, failed\n", id)
	}
}
func c42To32(total *int32, id int) {
	if atomic.CompareAndSwapInt32(total, 42, 32) {
		fmt.Printf("42 to 32:id %d, work\n", id)
	} else {
		fmt.Printf("42 to 32:id%d, failed\n", id)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	count := int32(32)

	for i := 0; i < 3; i++ {

		go c42To32(&count, i)
		go c32To42(&count, i)
	}
	time.Sleep(time.Second * 1)
}

// reference: https://github.com/astaxie/gopkg/blob/master/sync/atomic/compareandswapint32.md
