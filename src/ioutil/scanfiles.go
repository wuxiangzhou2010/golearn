package main

// 查找文件， 读一行

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// func main() {
// 	files, err := ioutil.ReadDir("../")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, f := range files {
// 		fmt.Println(f.Name())
// 	}
// }

// func main() {
// 	files, err := filepath.Glob("../*")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(files)
// }

func main() {
	var files []string
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Scan()
		fmt.Println(scanner.Text())

		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}
