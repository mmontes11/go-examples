package main

import (
	"log"
	"sync"
)

type counter struct {
	v   int
	mux sync.RWMutex
}

func (c *counter) increment() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v++
}

func (c *counter) value() int {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.v
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

	log.Printf("Counter: %d", c.value())
}
