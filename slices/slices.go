package main

import "log"

func printSlice(s []int) {
	log.Printf("%v len=%d cap=%d", s, len(s), cap(s))
}

func main() {
	log.Println("Nil slice")
	var s []int
	printSlice(s)

	log.Println("Literal slice")
	s = []int{1, 2, 3, 4, 5, 6}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3)
	printSlice(s)

	log.Println("Make slice with length")
	s = make([]int, 4)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = s[2:]
	printSlice(s)

	log.Println("Make slice with length and capacity")
	s = make([]int, 0, 5)
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = append(s, 1, 1)
	printSlice(s)

	s = s[1:]
	printSlice(s)

	s = s[1:7]
	printSlice(s)

	for i, v := range s {
		log.Printf("s[%d]=%d", i, v)
	}
}
