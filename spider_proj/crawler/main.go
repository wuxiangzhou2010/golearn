package main

import (
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/engine"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/scheduler"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/util/agent/my"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/zhenai/parser"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   scheduler.NewQueuedScheduler(),
		WorkerCount: 20,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
		Agent:      my.NewAgent(),
	})
}

func init() {
	// 方便通过web调试
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
}
