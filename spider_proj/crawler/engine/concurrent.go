package engine

import (
	"fmt"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}
type Scheduler interface {
	Submit(Request)
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		CreateWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	userLen := 0
	for {
		result := <-out

		for _, item := range result.Items {
			if v, ok := item.(model.Profile); ok {
				log.Printf("#%d: %v \n", userLen, v) // for profile
				userLen++
			} else {
				//log.Printf("Got item: %+v \n", item) // for city  and user name
			}
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}
func CreateWorker(out chan ParseResult, ready ReadyNotifier) {
	in := make(chan Request)
	go func() {

		for {
			//tell scheduler I'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				fmt.Printf("########Error %v Try to get %s \n", err, request.Url)
				continue
			}
			out <- result
		}
	}()
}
