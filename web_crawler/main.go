package main

import (
	"fmt"
	"sync"
)

const parallelism = 2

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type db struct {
	mu     sync.Mutex
	record map[string]int
}

func (d *db) Save(url string) {
	d.mu.Lock()

	d.record[url] += 1

	d.mu.Unlock()
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, db *db, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		db.Save(url)

		// fmt.Println(err)
		return
	}

	if _, ok := db.record[url]; ok {
		db.Save(url)

		// fmt.Printf("record %q exists\n", url)
		return
	}

	// fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		db.Save(url)

		go Crawl(u, depth-1, fetcher, db, wg)

		wg.Add(1)
	}
}

func main() {
	db := db{record: make(map[string]int)}
	wg := new(sync.WaitGroup)
	wg.Add(parallelism)

	for i := parallelism; i > 0; i-- {
		go Crawl("https://golang.org/", 4, fetcher, &db, wg)
	}

	wg.Wait()

	fmt.Printf("final map: %v\n", db.record)
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
