package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter can be used concurrently
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments a key
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

// Value gets value from a key
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	key := "key"
	for i := 0; i < 10; i++ {
		go c.Inc(key)
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value(key))
}
