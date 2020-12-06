package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func main() {
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func(wg *sync.WaitGroup) {
			atomic.AddUint64(&counter, 1)
			wg.Done()
		}(&wg)
		go func(wg *sync.WaitGroup) {
			atomic.AddUint64(&counter, 1)
			wg.Done()
		}(&wg)
	}

	wg.Wait()

	log.Printf("Counter: %d", counter)
}
