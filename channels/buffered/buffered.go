package main

import "log"

func main() {
	c := make(chan int, 3)

	log.Print("Sending 1, 2, 3...")
	c <- 1
	c <- 2
	c <- 3
	// Deadlock!
	// c <- 4

	log.Print("Received:")
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	// Deadlock!
	// log.Println(<-c)

	log.Print("Sending 4...")
	c <- 4

	log.Printf("Received %d", <-c)

	log.Print("Sending 5, 6, 7...")
	c <- 5
	c <- 6
	c <- 7

	log.Print("Received:")
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
}
