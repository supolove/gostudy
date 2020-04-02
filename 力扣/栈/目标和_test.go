package 栈

import (
	"fmt"
	"testing"
)

/*
目标和
https://leetcode-cn.com/explore/learn/card/queue-stack/219/stack-and-dfs/885/
*/

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 1 {
		if nums[0] == S || nums[0] == -S {
			return 1
		}
		return 0
	}
	result := 0

	var findTargetSumWaysDfs func(nums []int, S, index, sum int)
	findTargetSumWaysDfs = func(nums []int, S, index, sum int) {
		if index == len(nums) {
			if sum == S {
				result++
			}
			return
		}

		findTargetSumWaysDfs(nums, S, index+1, sum+nums[index])
		findTargetSumWaysDfs(nums, S, index+1, sum-nums[index])
	}

	findTargetSumWaysDfs(nums, S, 0, 0)
	return result
}

func Test_findTargetSumWays(t *testing.T) {
	a := []int{1, 2, 7, 9, 981}
	fmt.Println(findTargetSumWays(a, 1000000000))
}
