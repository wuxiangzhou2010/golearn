package main

import (
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/engine"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/util/agent/my"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
		Agent: my.NewAgent(),
	})
}
