package main

import (
	"log"
	"sync"
)

type counter struct {
	value int
	mux   sync.RWMutex
}

func (c *counter) increment() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value++
}

func (c *counter) getValue() int {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.value
}

func increment(counter *counter, wg *sync.WaitGroup) {
	defer wg.Done()

	counter.increment()
}

func main() {
	c := counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(&c, &wg)
		go increment(&c, &wg)
	}

	wg.Wait()

	log.Printf("Counter: %d", c.getValue())
}
