package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

const url = "https://xxx.freeimage.us/image.php?id=1B82_5CB6FEA4&jpg"

func main() {
	tr := &http.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: tr, //解决x509: certificate signed by unknown authority
	}
	req, err := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)

	if err != nil {
		panic(err)

	}
	if res.StatusCode != http.StatusOK {
		log.Println("Error: status code", res.StatusCode)
		panic(fmt.Errorf("error: status code", res.StatusCode))
	}
	defer res.Body.Close()
}
