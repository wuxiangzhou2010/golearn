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
	ConfigureWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.ConfigureWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		CreateWorker(in, out)

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
func CreateWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {

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
