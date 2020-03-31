package 队列

import (
	"fmt"
	"math"
	"testing"
)

/**
题目:
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:
	输入: n = 12
	输出: 3
	解释: 12 = 4 + 4 + 4.

示例 2:
	输入: n = 13
	输出: 2
	解释: 13 = 4 + 9.

使用贪心算法，并不合适
图解法

四平方和定理:
	 任何一个正整数都可以表示成不超过四个整数的平方之和。
	满足四数平方和定理的数n（这里要满足由四个数构成，小于四个不行），必定满足 n=4^(a(8b+7))

*/

// 四平方和定理解法
func numSquares(n int) int {

	//先根据上面提到的公式来缩小n
	for n%4 == 0 {
		n /= 4
	}

	//如果满足公式 则返回4
	if n%8 == 7 {
		return 4
	}

	//在判断缩小后的数是否可以由一个数的平方或者两个数平方的和组成
	var a int = 0
	for ; a*a <= n; a++ {
		b := int(math.Sqrt(float64(n - a*a)))
		if a*a+b*b == n {
			v := 0
			if a > 0 {
				v++
			}
			if b > 0 {
				v++
			}
			return v
		}
	}
	return 3
}

// 广度搜索算法
func numSquares2(n int) int {
	return 0
}

// 动态规划
/*
这道题类似于背包问题
物品：完全平方数，每种都有无数个
包的容积：n
目标：恰好把包放满时，物品的个数最少

动态规划的问题关键在于状态转移方程，这里的思路还是和前面使用图的广度优先遍历一样。例如
*/
func numSquares3(n int) int {
	return 0
}

func Test_numSquares(t *testing.T) {
	fmt.Println(numSquares(6))
}
