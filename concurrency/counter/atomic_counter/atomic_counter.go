package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func increment(counter *int32, wg *sync.WaitGroup) {
	defer wg.Done()

	atomic.AddInt32(counter, 1)
}

func main() {
	counter := int32(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(&counter, &wg)
		go increment(&counter, &wg)
	}

	wg.Wait()

	log.Printf("Counter: %d", atomic.LoadInt32(&counter))
}
