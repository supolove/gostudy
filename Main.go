package main

import (
	"fmt"
	"time"
)

var a string
var done chan int

func setup() {
	a = "hello, world"
	time.Sleep(time.Second * 5)
	done <- 1
}

var cccc chan int

func testBinxing4(i int) {
	go func() {
		d := <-cccc
		if i == d {
			fmt.Println("is true")
		} else {
			fmt.Println("is false")
		}

	}()
}

func tttt() {
	for i := 0; i < 10; i++ {
		testBinxing4(i)
	}
	cccc <- 1
	time.Sleep(time.Second)
	cccc <- 1
	for {

	}
}

// 返回生成自然数序列的管道
func generateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func primeFilter(ch <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			var i = <-ch
			//fmt.Println("计算：", i, prime)
			if i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

// 素数筛: 菊花链模型

func main() {
	var j = 0

	func() {
		for i := 0; i < 5; i++ {
			fmt.Println("j", j)
			i := j + 1
			j := i
			fmt.Println(j)
			j = i
		}

	}()
	//
	//for i := 0; i < 5; i++ {
	//	i := i
	//	func() {
	//		println(i)
	//	}()
	//}

	a, b := 2, 3
	max := If(a > b, a, b).(int)
	println(max)
}
