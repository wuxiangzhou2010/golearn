package chromedp

import "testing"

func TestAgent_Get(t *testing.T) {

		cdp := NewAgent()
		cdp.Get("http://www.zhenai.com/zhenghun/aba")

}