package parser

import (
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/engine"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/model"
	"regexp"
	"strconv"
)

var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3="">([^<]+)</h1>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([0-9]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var tallRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([0-9]+cm)</div>`)
var salaryRe = regexp.MustCompile(`<div class="m-btn" data-v-bff6f798="">月薪:([^<]+)</div>`)
var placeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">工作地:([^<]+)</div>`)

func ParseProfile(contents []byte) engine.ParseResult {

	profile := model.Profile{}

	profile.Name = extractString(contents, nameRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Height = extractString(contents, tallRe)
	profile.Income = extractString(contents, salaryRe)
	profile.Place = extractString(contents, placeRe)
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	result := engine.NewParseResult()
	result.Items = []interface{}{profile}
	return *result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
