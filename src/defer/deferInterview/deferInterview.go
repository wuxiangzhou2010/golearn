package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil { //  recover should be in the same goroutine when panics
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {

	for msg := range msgchan {
		m := msg.(int) // check type assertion
		fmt.Println("msg: ", m)

	}
}
func (p *Project) run(msgchan chan interface{}) {
	for {
		defer p.deferError()
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(1000000 * time.Second)
}

func main() {
	p := new(Project)
	p.Main()

}

/*
issue:
1. time.Sleep参数太大
2. 类型断言没有使用ok判断,类型不匹配
3. recover是在另一个goroutine中执行，当前goroutione不能捕获，导致直接panic

reference : https://zhuanlan.zhihu.com/p/35058068 issue 7
*/
