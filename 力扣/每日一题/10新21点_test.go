package 每日一题

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/problems/new-21-game/
新21点
动态规划
反序求解
*/

func minv(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}
func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K+W+1)
	for i := K; i <= N && i < K+W; i++ {
		dp[i] = 1.0
	}
	S := float64(minv(N-K+1, W))
	for i := K - 1; i >= 0; i-- {
		//for j:=1; j <= W ;j++  {
		//	dp[i] += dp[i + j] / float64(W)
		//}
		dp[i] = S / float64(W)
		S += dp[i] - dp[i+W]

	}
	return dp[0]
}

func Test_new21Game(t *testing.T) {
	fmt.Println(new21Game(21, 17, 10))
}
