package 递归

import (
	"fmt"
	"testing"
)

// 0 1 1 2 3 5 8

var cache = map[int]int{}

func fib(N int) int {
	n, ok := cache[N]
	if ok {
		return n
	}
	if N < 2 {
		cache[N] = N
		return N
	}

	cache[N] = fib(N-1) + fib(N-2)
	return cache[N]
}

func Test_fib(t *testing.T) {
	fmt.Println(fib(6))
}
