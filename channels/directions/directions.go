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
	pingChan := make(chan string, 1)
	pongChan := make(chan string, 1)

	ping(pingChan, "Hello world!")
	pong(pingChan, pongChan)

	log.Println(<-pongChan)
}
