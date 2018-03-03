package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

type EmployeeIdGenerator func(company string, department string, sn uint32) string

var c = "Gophers"

var sn uint32

func generateId(generator EmployeeIdGenerator, department string) (string, bool) {
	if generator == nil {
		fmt.Println("1")
		return "false", false
	}
	newSn := atomic.AddUint32(&sn, 1)
	fmt.Println("2")
	return generator(c, department, newSn), true
}

func appendSn(firstPart string, sn uint32) string {
	return firstPart + strconv.FormatUint(uint64(sn), 10)
}

func main() {

	var generator EmployeeIdGenerator
	generator = func(company string, department string, sn uint32) string {
		return appendSn(c+department, sn)
	}
	fmt.Println(generateId(generator, "RD"))
}
