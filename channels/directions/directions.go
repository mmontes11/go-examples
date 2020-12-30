package main

import "log"

func ping(c chan<- string, msg string) {
	c <- msg
}

func pong(source <-chan string, target chan<- string) {
	msg := <-source
	target <- msg
}

func main() {
	pingChan := make(chan string)
	pongChan := make(chan string)

	go pong(pingChan, pongChan)
	// Deadlock!
	// pong(pingChan, pongChan)
	ping(pingChan, "Hello unbuffered world!")

	log.Println(<-pongChan)

	pingChan = make(chan string, 1)
	pongChan = make(chan string, 1)

	ping(pingChan, "Hello buffered world!")
	pong(pingChan, pongChan)

	log.Println(<-pongChan)
}
