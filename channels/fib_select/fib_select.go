package main

import "fmt"

func fibonacci(c chan int, quit chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- struct{}{}
	}()
	fibonacci(c, quit)
}
