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

type valueOp struct {
	op
}

func newIncrementOp() incrementOp {
	return incrementOp{
		op: op{
			res: make(chan int),
		},
	}
}

func newValueOp() valueOp {
	return valueOp{
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
	valueOps := make(chan valueOp)
	go func() {
		val := 0
		for {
			select {
			case op := <-incrementOps:
				val++
				op.res <- val
			case op := <-valueOps:
				op.res <- val
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

	valueOp := newValueOp()
	valueOps <- valueOp

	log.Printf("Counter: %d", <-valueOp.res)
}
