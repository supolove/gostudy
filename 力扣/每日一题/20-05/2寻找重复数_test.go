package _0_05

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/problems/find-the-duplicate-number/

2020-05-26
*/

func findDuplicate(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v]++
	}
	return 0
}

// 环的思想解题
// 重复数字是环的入口
// 快指针和慢指针终究会相遇
func findDuplicate2(nums []int) int {
	slow := nums[0]
	fast := nums[slow]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = nums[slow]
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func Test_findDuplicate(t *testing.T) {
	fmt.Println(findDuplicate2([]int{1, 3, 4, 2, 2}))
}
