package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker #%v is pulling off job = %v\n", id, job)
		time.Sleep(2 * time.Second)
		results <- job
	}
	time.Sleep(5 * time.Second)
	close(results)
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	//start workers routine, this should give us 3 go routines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//send jobs to channel so workers have work
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= 10; i++ {
		fmt.Println("print results = ", <-results)
	}

}
