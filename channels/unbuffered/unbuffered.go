package main

import "log"

func main() {
	c := make(chan int)

	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		log.Printf("Sending %d...", v)
		go func(v int) {
			c <- v
		}(v)
		log.Printf("Received %d", <-c)
	}
}
