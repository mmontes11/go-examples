package main

import (
	"fmt"
	"sync"
	"time"
)

// Crawled represents the number of fetches of a URL in a concurrent safe way
type Crawled struct {
	urls map[string]int
	mux  sync.Mutex
}

// Add marks a URL as crawled
func (c *Crawled) Add(url string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.urls[url]++
}

// IsCrawled checks if a URL is crawled
func (c *Crawled) IsCrawled(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.urls[url] > 0
}

// Fetcher retrieves URLs
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, crawled *Crawled) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if crawled.IsCrawled(url) {
		return
	}
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	crawled.Add(url)
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, crawled)
	}
	return
}

func main() {
	crawled := Crawled{urls: make(map[string]int)}
	go Crawl("https://golang.org/", 4, fetcher, &crawled)
	time.Sleep(time.Second)
	fmt.Println("Crawler stats:")
	for url, count := range crawled.urls {
		fmt.Printf("%v: %d\n", url, count)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
