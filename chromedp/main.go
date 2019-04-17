// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string

	err := chromedp.Run(ctx,
		// 访问页面
		chromedp.Navigate(`http://www.zhenai.com/zhenghun/aba`),
		// 等待列表渲染
		chromedp.Sleep(5*time.Second),
		// 获取获取服务列表HTML
		chromedp.OuterHTML("#app", &example, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
