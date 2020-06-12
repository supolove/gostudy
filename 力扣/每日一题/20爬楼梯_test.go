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

// 动态规划
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 滚动数组
func climbStairs3(n int) int {
	p, q, r := 0, 0, 1
	for i := 0; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func Test(t *testing.T) {
	fmt.Println(climbStairs3(5))
}
