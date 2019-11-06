package main

import (
	"fmt"
	"gostudy/basetest"
	"io"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

//func main() {
//	fmt.Println("hello go")

/*
	//第二个参数 chars 中任意一个字符（Unicode Code Point）如果在第一个参数 s 中存在，则返回true。
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("ins failure", "sg"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
*/
/*
	子串出现次数(字符串匹配)
	朴素匹配算法
	KMP算法
	Rabin-Karp算法
	Boyer-Moore算法

	charStr := "我的abcde"
	nameRune := []rune(charStr)
	fmt.Println(string(nameRune[:3]))
*/

//fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

/*
	var strArr []string
	strArr = []string{"a","c"}

	fmt.Println(strings.Join(strArr, "b"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
*/
//strings.NewReplacer("")
//testFibonacci()
//testFibonacci3()
//}

// 斐波那契数列
func testFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	x, y := 1, 0
	return func() int {
		x, y = y, x+y
		return x
	}
}

func fibonacci2(n int, c chan int) {
	x, y := 1, 0
	for i := 0; i < n; i++ {
		x, y = y, x+y
		c <- x
	}
	close(c)
}

func testFibonacci2() {
	c := make(chan int, 10)
	go fibonacci2(cap(c), c)
	for v := range c {
		fmt.Println(v)
	}
}

func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return

		}
	}
}

func testFibonacci3() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
			time.Sleep(time.Second)
		}
		quit <- 0
	}()

	fibonacci3(c, quit)
}

// rot13Reader字符加密

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) testRot13Reader(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	if err == io.EOF {
		return 0, io.EOF
	}

	for i, c := range b {
		switch {
		case c < 'N':
			b[i] = c + 13
			fallthrough
		case c <= 'Z':
			b[i] = c - 13
		case c < 'n':
			b[i] = c + 13
		case c <= 'z':
			b[i] = c - 13
		default:
			b[i] = c
		}
	}
	return n, nil
}

/*
比较二叉平衡树内容是否相同
这道题可以练习一下二叉树的遍历、channel的状态测试等知识点
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkInternal(t, ch)
	close(ch)
}

func WalkInternal(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkInternal(t.Left, ch)
		ch <- t.Value
		WalkInternal(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if ok1 == false {
			break
		}
	}

	return true
}

func testTree() {
	//ch := make(chan int)
	//go Walk(tree.New(2), ch)
	//for i := range ch{
	//	fmt.Println(i)
	//}
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

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

/////
func testPachong() {
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(2000 * time.Millisecond)
}

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

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
}

func main() {
	basetest.PrintTest()
}

// func main() {
// 	basetest.PrintTest()
// }
