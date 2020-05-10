package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, value := range strings.Fields(s) {
		m[value]++
	}
	return m
}

func main() {
	wc.Test(wordCount)
}
