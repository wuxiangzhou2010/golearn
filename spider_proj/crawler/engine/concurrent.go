package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}
type Scheduler interface {
	Submit(*Request)
	ConfigureWorkerChan(chan *Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan *Request)
	out := make(chan *ParseResult)

	e.Scheduler.ConfigureWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createworker(in, out)

	}

	for _, r := range seeds {
		e.Scheduler.Submit(&r)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			fmt.Printf("Got item: %v \n", item)
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(&request)
		}
	}
}
func createworker(in chan *Request, out chan *ParseResult) {
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
