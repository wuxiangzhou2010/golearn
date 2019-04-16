package parser

import (
	"fmt"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/engine"
	"regexp"
)

//<a href="http://album.zhenai.com/u/1610362984" target="_blank">青瓷雪</a>
//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
//var cityRe = regexp.MustCompile(`<a href="http://album.zhenai.com/u/[0-9^>]+" target="_blank"[^>]*>([^<]+)</a>`)
var cityRe = regexp.MustCompile(`http://album.zhenai.com`)

//const cityRe = `(album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	//re := regexp.MustCompile(cityRe)
	fmt.Println("ParseCity")
	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseProfile,
		})

	}
	return result
}
