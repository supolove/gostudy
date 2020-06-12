package 每日一题

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/problems/climbing-stairs/
*/

func digui2(n int, m map[int]int) int {
	v, ok := m[n]
	if ok {
		return v
	}
	if n == 1 {
		m[1] = 1
		return 1
	}

	if n == 2 {
		m[2] = 2
		return 2
	}
	res := digui2(n-1, m) + digui2(n-2, m)
	m[n] = res
	return res
}

func climbStairs(n int) int {
	return digui2(n, map[int]int{})
}

func Test(t *testing.T) {
	fmt.Println(climbStairs(5))
}
