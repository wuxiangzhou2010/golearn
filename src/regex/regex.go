package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com@abc.com
email1 is  abc@def.org
email2 : 2323@qqq.com
email3 is dd@erer.ere.cn`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+)(\.[a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	// fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}

}
