package main

import "log"

func main() {
	c := make(chan int, 3)

	log.Print("Sending 1, 2, 3...")
	for _, v := range []int{1, 2, 3} {
		go func(val int) {
			c <- val
		}(v)
	}
	log.Print("Sending 4")
	c <- 4

	log.Print("Received:")
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	log.Println(<-c)
	// Deadlock!
	// log.Println(<-c)

	values := []int{5, 6, 7, 8, 9, 10}
	log.Printf("Sending %v...", values)
	for _, v := range values {
		// Deadlock!
		// c <- v
		go func(val int) {
			c <- val
		}(v)
	}

	log.Print("Received:")
	for range values {
		log.Println(<-c)
	}
}
