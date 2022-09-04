package main

import "log"

func printSlice(s []int) {
	log.Printf("%v len=%d cap=%d", s, len(s), cap(s))
}

func main() {
	log.Println("Nil slice")
	var s []int
	printSlice(s) // [] len=0 cap=0

	log.Println("Literal slice")
	s = []int{1, 2, 3, 4, 5, 6}
	printSlice(s) // [1 2 3 4 5 6] len=6 cap=6

	s = s[:0]
	printSlice(s) // [] len=0 cap=6

	s = s[:4]
	printSlice(s) // [1 2 3 4] len=4 cap=6

	s = s[2:]
	printSlice(s) // [3 4] len=2 cap=4

	s = s[2:]
	printSlice(s) // [] len=0 cap=2

	s = append(s, 1)
	printSlice(s) // [1] len=1 cap=2

	s = append(s, 2, 3)
	printSlice(s) // [1 2 3] len=3 cap=4

	log.Println("Make slice with length")
	s = make([]int, 4)
	printSlice(s) // [0 0 0 0] len=4 cap=4

	s = append(s, 1)
	printSlice(s) // [0 0 0 0 1] len=5 cap=8

	s = s[2:]
	printSlice(s) // [0 0 1] len=3 cap=6

	log.Println("Make slice with length and capacity")
	s = make([]int, 0, 5)
	printSlice(s) // [] len=0 cap=5

	s = s[:4]
	printSlice(s) // [0 0 0 0] len=4 cap=5

	s = append(s, 1, 1)
	printSlice(s) // [0 0 0 0 1 1] len=6 cap=10

	s = s[1:]
	printSlice(s) // [0 0 0 1 1] len=5 cap=9

	s = s[1:7]
	printSlice(s) // [0 0 1 1 0 0] len=6 cap=8

	// s[0]=0
	// s[1]=0
	// s[2]=1
	// s[3]=1
	// s[4]=0
	// s[5]=0
	for i, v := range s {
		log.Printf("s[%d]=%d", i, v)
	}
}
