package algorithm

import (
	"fmt"
	"testing"
	"time"
)

/*
	斐波那契数列：1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144,
*/

// 斐波那契数列
func testFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	x, y := 0, 1
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
	go fibonacci2(5, c)
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

func TestFibonacci(t *testing.T) {
	testFibonacci2()
}
