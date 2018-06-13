package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://baidu.com",
		"http://golang.org",
		"http://amazon.com",
		"http://jd.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		// fmt.Println(<-c)
		go func(s string) {
			time.Sleep(3 * time.Second)
			checkLink(s, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
