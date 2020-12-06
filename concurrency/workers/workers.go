package main

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type state struct {
	mux  sync.RWMutex
	data int
}

type latestN struct {
	mux  sync.RWMutex
	N    int
	data []int
}

func (l *latestN) add(i int) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if len(l.data) == l.N {
		l.data = l.data[1:]
	}
	l.data = append(l.data, i)
}

func newLatestN(n int) latestN {
	return latestN{
		N:    n,
		data: make([]int, n),
	}
}

func printStats(s *state, l *latestN, stateOps, lastNops *uint64) {
	s.mux.RLock()
	l.mux.RLock()
	defer func() {
		s.mux.RUnlock()
		l.mux.RUnlock()
	}()

	log.Println("Stats 📊")
	log.Printf("state: %v", s.data)
	log.Printf("last%d: %v", l.N, l.data)
	log.Printf("state ops: %d", atomic.LoadUint64(stateOps))
	log.Printf("last%d ops: %d", l.N, atomic.LoadUint64(lastNops))
}

func main() {
	var lastNops uint64
	var stateOps uint64

	s := state{}
	l := newLatestN(10)

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		wg.Add(1)
		time.Sleep(100 * time.Millisecond)

		go func(wg *sync.WaitGroup, ctx context.Context) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					log.Print("Finishing lastN worker...")
					return
				default:
					s.mux.RLock()
					l.add(s.data)
					s.mux.RUnlock()

					atomic.AddUint64(&lastNops, 1)

					time.Sleep(100 * time.Millisecond)
				}
			}
		}(&wg, ctx)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ctx context.Context) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					log.Print("Finishing state worker...")
					return
				default:
					s.mux.Lock()
					s.data++
					s.mux.Unlock()

					atomic.AddUint64(&stateOps, 1)
				}
			}
		}(&wg, ctx)
	}

	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		printStats(&s, &l, &stateOps, &lastNops)
	}

	cancel()
	wg.Wait()
}
