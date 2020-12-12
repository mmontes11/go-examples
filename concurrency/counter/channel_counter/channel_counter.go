package main

import (
	"log"
	"sync"
)

type op struct {
	res chan int
}

type incrementOp struct {
	op
}

type getValueOp struct {
	op
}

func newIncrementOp() incrementOp {
	return incrementOp{
		op: op{
			res: make(chan int),
		},
	}
}

func newGetValueOp() getValueOp {
	return getValueOp{
		op: op{
			res: make(chan int),
		},
	}
}

func increment(ops chan<- incrementOp, wg *sync.WaitGroup) {
	defer wg.Done()

	op := newIncrementOp()
	ops <- op
	<-op.res
}

func main() {
	incrementOps := make(chan incrementOp)
	getValueOps := make(chan getValueOp)

	go func() {
		counter := 0
		for {
			select {
			case op := <-incrementOps:
				counter++
				op.res <- counter
			case op := <-getValueOps:
				op.res <- counter
			}
		}
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(incrementOps, &wg)
		go increment(incrementOps, &wg)
	}

	wg.Wait()

	getValueOp := newGetValueOp()
	getValueOps <- getValueOp

	log.Printf("Counter: %d", <-getValueOp.res)
}
