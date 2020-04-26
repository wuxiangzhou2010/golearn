package main

import "fmt"

func worker(id int, work chan int, result chan int) {

	for {
		if v, ok := <-work; ok {
			fmt.Printf("worker: %d, get work: %d", id, v)
			result <- v * 2
		} else {
			break
		}
	}
}

func main() {

	workChan := make(chan int, 1000)
	resultChan := make(chan int, 10)
	
	for i := 0; i < 3; i++ {
		go worker(i, workChan, resultChan)
	}

	for i := 0; i < 10; i++ {
		workChan <- i
	}
	close(workChan)

	for i := 0; i < 10; i++ {
		fmt.Println("main get result: ", <-resultChan)
	}

}
