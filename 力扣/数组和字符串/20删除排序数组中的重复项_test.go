package 数组和字符串

import (
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/202/conclusion/795/
*/

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			break
		}
		if nums[i] == nums[i+1] {
			for j := i; j < len(nums); j++ {
				if j == len(nums)-1 {
					break
				}
				nums[j] = nums[j+1]
			}
			nums = nums[:len(nums)-1]
			i--
		}
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	s := 0
	for f := 1; f < len(nums); f++ {
		if nums[f] != nums[s] {
			s++
			nums[s] = nums[f]
		}
	}
	nums = nums[:s+2]
	return s + 1
}

func Test_removeDuplicates(t *testing.T) {
	removeDuplicates([]int{1, 1, 1, 2})
}
