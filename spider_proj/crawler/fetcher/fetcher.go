package fetcher

import (
	"bufio"
	"github.com/wuxiangzhou2010/daily_learning/go/spider_proj/crawler/util/agent/my"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type Fetcher interface {
	Fetch(url string) ([]byte, error)
}

func Fetch(url string) ([]byte, error) {
	res, err := my.NewAgent().Get("http://www.zhenai.com/zhenghun")
	if err != nil {

		return nil, err
	}
	//if res.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code", res.StatusCode)
	//	return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
	//}
	//defer res.Body.Close()

	bodyReader := bufio.NewReader(res)
	e := determinEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

func determinEncoding(r *bufio.Reader) encoding.Encoding {

	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
