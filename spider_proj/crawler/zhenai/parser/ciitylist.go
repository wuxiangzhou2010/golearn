package parser

import (
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/engine"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/util/agent/chromedp"

	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte) *engine.ParseResult {

	//re := regexp.MustCompile(cityListRe)
	matches := cityListRe.FindAllSubmatch(contents, -1)
	result := engine.NewParseResult()
	for _, m := range matches {
		//result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
			Agent: chromedp.NewAgent(),
			Name: "City: " + string(m[2]),
		})

	}
	return result

}
