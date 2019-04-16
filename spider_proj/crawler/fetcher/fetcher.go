package fetcher

import (
	"encoding"
	"fmt"
	//"golang.org/x/text/encoding"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
)

type Fetcher interface {
	Fetch(url string) ([]byte, error)
}

func Fetch(url string) ([]byte, error) {
	res, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		//panic(err)
		return nil, err
	}

	defer res.Body.Close()
	//charset.De
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", res.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}
	return ioutil.ReadAll(res.Body)

}

func determinEncoding(r io.Reader) encoding.Encoding {

	charset.DeterminEncoding()
}
