package parser

import (
	"github.com/wuxiangzhou2010/luandun/go/spider_proj/crawler/engine"
	"github.com/wuxiangzhou2010/luandun/go/spider_proj/crawler/util/agent/chromedp"
	"github.com/wuxiangzhou2010/luandun/go/spider_proj/crawler/util/str"

	"log"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	contents = []byte(str.DelLongSlash(string(contents)))
	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.NewParseResult()
	log.Println("ParseCity total matches ", len(matches))
	for _, m := range matches {

		result.Requests = append(result.Requests, engine.Request{ // what need to be process again
			Url:        string(m[1]),
			ParserFunc: ParseProfile,
			Agent:      chromedp.NewAgent(), // data exist in dynamic area
			Name:       string(m[2]),
		})
		result.Items = append(result.Items, string(m[2])) //name

	}

	return *result
}
