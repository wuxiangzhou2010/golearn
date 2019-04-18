package scheduler

import "github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func NewQueuedScheduler() *QueuedScheduler {
	return &QueuedScheduler{}
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w

}

//func (s *QueuedScheduler) ConfigureWorkerChan(chan engine.Request) {
//	panic("implement me")
//}
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				//send r to ? worker
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				//send ? next_request to w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: //同时有request 和 worker
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()

}
