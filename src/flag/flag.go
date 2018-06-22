package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "Sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleep for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
