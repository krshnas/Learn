package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	for i := 0; i < 40; i++ {
		jobs <- i
	}

	close(jobs)
	for i := 0; i < 40; i++ {
		fmt.Println(<-results)
	}
	fmt.Println("execution time", time.Since(now).Seconds())
}

func worker(jobs <-chan int, results chan<- int) { // jobs will only receive and results will only sends channel, if we try to send on jobs then we will get compile time error
	for n := range jobs {
		results <- fib(n)
	}

}

func fib(num int) int {
	if num <= 1 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}
