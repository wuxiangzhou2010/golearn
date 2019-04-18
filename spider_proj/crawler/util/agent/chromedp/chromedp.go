package chromedp

import (
	"context"
	"io"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

//Agent is a chromedp agent
type Agent struct{}

func NewAgent() *Agent {
	return &Agent{}
}

func (c *Agent) Get(url string) (io.Reader, error) {
	//t1 := time.Now()
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
		//chromedp.Sleep(5*time.Second),
		// 获取获取服务列表HTML
		chromedp.WaitVisible(`#app`),
		chromedp.OuterHTML("#app", &example, chromedp.ByID),
	)
	//t2 := time.Now()

	//log.Println("time used: ", t2.Sub(t1), " to get ", url)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Get contents: %s", example)
	sr := strings.NewReader(example)
	return sr, nil
}
