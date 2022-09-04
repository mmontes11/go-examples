package main

import (
	"log"
	"time"
)

func worker(done chan struct{}) {
	log.Print("Working...")
	time.Sleep(time.Second)
	log.Print("Done!")

	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go worker(done)

	// If removed, program finish before worker starts
	<-done
}
