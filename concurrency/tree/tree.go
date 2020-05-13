package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch, quit)
	select {
	// Value successfully sent.
	case ch <- t.Value:
	case <-quit:
		return
	}
	walkImpl(t.Right, ch, quit)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch, quit chan int) {
	walkImpl(t, ch, quit)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	quit := make(chan int)
	defer close(quit)
	go Walk(t1, c1, quit)
	go Walk(t2, c2, quit)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	t1 := tree.New(10)
	c := make(chan int)
	quit := make(chan int)
	go Walk(t1, c, quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	t2 := tree.New(10)
	t3 := tree.New(20)
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t2, t3))
}
