package main

import "fmt"

func defered() string {
	defer fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	return "done"
}

func main() {
	fmt.Println(defered())
}
