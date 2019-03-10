package main

// 1. os.Getwd
// 2. path/filepath.Join
import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	if workingDIr, err := os.Getwd(); err == nil {
		fmt.Println(filepath.Join(workingDIr, "config.json"))
	}
}
