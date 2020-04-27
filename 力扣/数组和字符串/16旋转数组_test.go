package 数组和字符串

import (
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/202/conclusion/791/
*/

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	n := make([]int, 0)
	k = k % len(nums)
	for i := 0; i < len(nums); i++ {
		n = append(n, nums[i])
	}

	o := n[:len(nums)-k]
	p := n[len(nums)-k:]

	for i := 0; i < len(nums); i++ {
		if i < len(p) {
			nums[i] = p[i]
		} else {
			nums[i] = o[i-len(p)]
		}
	}
}

// 三次反转做法

func rotate2(nums []int, k int) {

	reverse := func(nums []int, start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	k %= len(nums)
	// 第一反转
	reverse(nums, 0, len(nums)-1)
	// 第二次
	reverse(nums, 0, k-1)
	// 第三次
	reverse(nums, k, len(nums)-1)
}

func Test_rotate(t *testing.T) {
	rotate([]int{1, 2, 3, 4, 5}, 3)
	/*
		5 2 3 4 1
		5 4 3 2 1
		4 5 3 2 1
		4 5 1 2 3
	*/
}
