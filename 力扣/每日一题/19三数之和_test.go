package 每日一题

/*
https://leetcode-cn.com/problems/3sum/
*/

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	return [][]int{}
}

func Test_threeSum(t *testing.T) {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}
