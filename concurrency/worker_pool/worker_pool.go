package main

import (
	"fmt"
	"log"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Printf("Worker %d started job %d", id, j)
		time.Sleep(time.Second)
		results <- j * 2
		log.Printf("Worker %d finished job %d", id, j)
	}
}

// time go run worker_pool.go
func main() {
	const numJobs = 5
	const numWorkers = 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 0; w < numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	log.Println("Results:")
	for r := 0; r < numJobs; r++ {
		fmt.Println(<-results)
	}
}
