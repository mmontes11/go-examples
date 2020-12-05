package main

import (
	"log"
	"time"
)

func send(c chan<- string, msg string, delay time.Duration) {
	time.Sleep(delay)
	c <- msg
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go send(c1, "Foo", 1*time.Second)
	go send(c2, "Bar", 2*time.Second)

	for i := 0; i < 2; i++ {
		select {
		case v1 := <-c1:
			log.Printf("Received from c1: %s", v1)
		case v2 := <-c2:
			log.Printf("Received from c2: %s", v2)
		}
	}

	go send(c1, "Foo", 1*time.Second)

	select {
	case v1 := <-c1:
		log.Printf("Received from c1: %s", v1)
	case <-time.After(5 * time.Second):
		log.Println("c1 timeout")
	}

	go send(c1, "Foo", 5*time.Second)

	select {
	case v1 := <-c1:
		log.Printf("Received from c1: %s", v1)
	case <-time.After(1 * time.Second):
		log.Println("c1 timeout")
	}

	select {
	case v1 := <-c1:
		log.Printf("Received from c1: %s", v1)
	default:
		log.Println("No messages in c1")
	}

	msg := "Foo"
	select {
	case c1 <- "Foo":
		log.Printf("Message \"%s\" sent to c1", msg)
	default:
		log.Printf("Unable to send \"%s\" to c1", msg)
	}

	go send(c1, "Foo", 0)
	select {
	case v1 := <-c1:
		log.Printf("Received from c1: %s", v1)
	default:
		log.Println("No messages in c1")
	}
}
