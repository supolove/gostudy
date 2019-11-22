package basetest

import (
	"fmt"
	"testing"
)

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func TestYouquSanyuan(t *testing.T) {
	// 三元表达式
	a, b := 2, 3
	max := If(a > b, a, b).(int)
	println(max)
}

func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			}
		}
	}()
	return c
}

// 基于管道随机数
func TestYouquSuiji(t *testing.T) {
	for i := range random(100) {
		fmt.Println(i)
	}
}
