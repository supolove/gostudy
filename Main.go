package main

import "fmt"

import "time"

import "sync"

var a string
var done chan int

func setup() {
	a = "hello, world"
	time.Sleep(time.Second * 5)
	done <- 1
}

// 用chan实现同步

func main() {

	fmt.Println("start main")

	done = make(chan int)

	go setup()
	<-done

	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("您好，世界")
		mu.Unlock()
	}()

	fmt.Println(a)
}
