package my

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Agent struct{}

func NewAgent() *Agent {
	return &Agent{}
}

func (m *Agent) Get(url string) (io.Reader, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close() // close reader closer

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	br := bytes.NewReader(bs)
	return br, nil

}
