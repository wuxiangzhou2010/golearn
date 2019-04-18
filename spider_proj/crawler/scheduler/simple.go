package scheduler

import "github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func NewSimpleScheduler() *SimpleScheduler {
	return &SimpleScheduler{}
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func(req engine.Request) {
		s.workerChan <- req
	}(r)
}
