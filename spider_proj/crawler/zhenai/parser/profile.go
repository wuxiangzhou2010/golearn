package parser

import (
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/engine"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/model"
	"regexp"
	"strconv"
)

var nameRe = regexp.MustCompile(`<span class="nickName" data-v-3c42fade="">([^<]+)</span>`)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

func ParseProfile(contents []byte) *engine.ParseResult {

	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRe)
	result := engine.NewParseResult()
	result.Items= []interface{}{profile}
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
