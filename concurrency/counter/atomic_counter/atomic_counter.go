package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func increment(counter *uint64, wg *sync.WaitGroup) {
	defer wg.Done()

	atomic.AddUint64(counter, 1)
}

func main() {
	counter := uint64(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(&counter, &wg)
		go increment(&counter, &wg)
	}

	wg.Wait()

	log.Printf("Counter: %d", counter)
}
