package _0_06

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode-cn.com/problems/best-sightseeing-pair/
*/

// 枚举，肯定会超时的
func maxScoreSightseeingPair(A []int) int {
	max := 0
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			v := A[i] + A[j] + i - j
			if v > max {
				max = v
			}
		}
	}
	return max
}

/*
可以先从暴力法中找思路
暴力法：算出每两对的 得分取最大值
        for (j : 1 -> A.length) {
            for (i : 0 -> j) {
                maxvalue = Math.max(maxvalue, A[j] - j   +   A[i] + i);
           }
        }
当然是超时的 -------> 优化 贪心算法
从暴力法可以看出在每次计算过程中 在每个 j 都固定 A[j] -j 计算它前面的 A[i] + i 的最大值 (i : 1 -> j)

贪心算法：将问题分解为 A[i] + i 和 A[j] - j 的最大值  i < j
A[j] - j是固定的，那我们在一次遍历 j 的时候只需要不断保存并且更新 A[i] + i 的值即可求出最大值。
*/
func maxScoreSightseeingPair2(A []int) int {
	left := A[0]
	res := math.MinInt32
	for i := 1; i < len(A); i++ {
		res = max(res, left+A[i]-i)
		left = max(left, A[i]+i)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_maxScoreSightseeingPair(t *testing.T) {
	fmt.Println(maxScoreSightseeingPair2([]int{8, 1, 7, 2, 6}))
}
