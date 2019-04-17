package engine

import (
	"io/ioutil"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s\n", r.Url)

		result:= worker(r)
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			log.Printf("Got item %s", item)
		}

	}
}


func worker(r Request) *ParseResult{
	re, err := r.Agent.Get(r.Url)
	if err != nil {
		panic(err)
	}
	all, err := ioutil.ReadAll(re)
	if err != nil {
		panic(err)
	}
	return  r.ParserFunc(all)

}