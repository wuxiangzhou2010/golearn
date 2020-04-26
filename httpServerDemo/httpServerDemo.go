package main

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, how are you")
}

func main() {
	http.HandleFunc("/", SayHello)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("http server failed: ", err)
	}
}