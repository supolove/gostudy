package study

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
4. 网络爬虫性能优化

a) 每个URL只需要被爬取一次

b) 并行爬取多个URL
*/
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	if cache.has(url) {
		fmt.Println("Hit cache")
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("found: %s %q\n", url, body)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}

	return
}

type Cache struct {
	cache map[string]bool
	mutex sync.Mutex
}

func (cache *Cache) add(url string) {
	cache.mutex.Lock()
	cache.cache[url] = true
	cache.mutex.Unlock()
}

func (cache *Cache) has(url string) bool {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	_, ok := cache.cache[url]
	if !ok {
		cache.cache[url] = true
	}
	return ok
}

var cache Cache = Cache{
	cache: make(map[string]bool),
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

/////

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
}

func TestPachong(t *testing.T) {
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(2000 * time.Millisecond)
}
