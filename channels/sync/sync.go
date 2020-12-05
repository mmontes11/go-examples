package main

import (
	"log"
	"time"
)

func worker(done chan bool) {
	log.Print("Working...")
	time.Sleep(time.Second)
	log.Print("Done!")

	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)

	// If removed, program finish before worker starts
	<-done
}
