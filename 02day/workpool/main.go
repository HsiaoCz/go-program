package main

import (
	"fmt"
	"time"
)

// work pool
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启goroutine
	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}
	// 开启五个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	close(results)
}

func worker(id int, job <-chan int, result chan<- int) {
	for j := range job {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		result <- j * 2
	}
}
