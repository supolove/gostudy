package 递归

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/orignial/card/recursion-i/259/complexity-analysis/1227/
*/

func digui6(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 || n == -1 {
		return x
	}

	if n%2 > 0 {
		return x * digui6(x, n-1)
	}
	if n%2 < 0 {
		return x * digui6(x, n+1)
	}
	a := digui6(x, n/2)
	return a * a

}

func myPow(x float64, n int) float64 {
	r := digui6(x, n)
	if n < 0 {
		r = 1 / r
	}
	return r
}

func Test_myPow(t *testing.T) {
	//0.00001
	//2147483647
	fmt.Println(myPow(2, -5))
}
