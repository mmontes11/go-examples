package main

import "log"

func main() {
	jobs := make(chan int, 5)
	done := make(chan struct{})

	go func() {
		for {
			j, more := <-jobs
			if more {
				log.Printf("Received job %d", j)
			} else {
				log.Print("No more jobs")
				done <- struct{}{}
				return
			}
		}
	}()
	for i := 0; i < 10; i++ {
		log.Printf("Sending job %d...", i)
		jobs <- i
	}
	close(jobs)

	<-done

	jobs = make(chan int, 5)
	log.Println()

	go func() {
		for i := 0; i < 10; i++ {
			log.Printf("Sending job %d...", i)
			jobs <- i
		}
		close(jobs)
	}()
	for j := range jobs {
		log.Printf("Received job %d...", j)
	}
}
