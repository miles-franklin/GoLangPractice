// Dropping this for now, since the 

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Cache struct {
	urls map[string]string
	mu sync.Mutex
}

func (c Cache) Add(key string, value string) (){
	c.mu.Lock()
	c.urls[key] = value
	c.mu.Unlock()
}

func (c Cache) Check(key string) (string, bool){
	c.mu.Lock()
	val, ok := c.urls[key]
	defer c.mu.Unlock()

	return val, ok
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *Cache, ch chan string){
	defer close(ch)

	if depth <= 0 {
		return
	}

	// TODO: Don't fetch the same URL twice. [DONE]
	_, ok := (*cache).Check(url)

	if !ok {
		body, urls, err := fetcher.Fetch(url)

		(*cache).Add(url, body)

		if err != nil {
			ch <- err.Error()
			return
		}
		ch <- fmt.Sprintf("found: %s %q\n", url, body)

		// TODO: Fetch URLs in parallel.
		// "every instance of Crawl gets its own return channel and the caller
		// function collects the results in its return channel."
		// https://stackoverflow.com/questions/13217547/tour-of-go-exercise-10-crawler
		channels := make([]chan string, len(urls))
		for i, u := range urls {
			channels[i] = make(chan string)
			go Crawl(u, depth-1, fetcher, cache, channels[i])
		}

		for i := range channels {
			for val := range channels[i] {
				ch <- val
			}
		}

	}
}

func main() {
	cache := Cache{}
	cache.urls = map[string]string{}

	ch := make(chan string)

	go Crawl("https://golang.org/", 4, fetcher, &cache, ch)

	for val := range ch {
		fmt.Printf("%v", val)
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
