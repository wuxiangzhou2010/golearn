package chromedp

import (
	"context"
	"github.com/chromedp/chromedp"
	"io"
	"log"
	"strings"
	"time"
)

type Agent struct {
}

func NewAgent() *Agent {
	return &Agent{}
}

func (c *Agent) Get(url string) ( io.Reader, error) {

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
		chromedp.Navigate(url),
		// 等待列表渲染
		chromedp.Sleep(5*time.Second),
		// 获取获取服务列表HTML
		chromedp.OuterHTML("#app", &example, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Get contents: %s", example)
	sr:= strings.NewReader(example)
	return sr, nil
}

