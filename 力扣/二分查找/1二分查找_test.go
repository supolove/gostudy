package 二分查找

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/binary-search/208/background/833/
*/

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for {
		mid := (right-left)/2 + left

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if right == mid {
				if nums[left] == target {
					return left
				} else {
					return -1
				}
			}
			right = mid
		} else {
			if left == mid {
				if nums[right] == target {
					return right
				} else {
					return -1
				}
			}
			left = mid
		}
	}
}

func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// 获取中间位置
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 中间+1
			left = mid + 1
		} else {
			right = right - 1
		}
	}
	return -1
}

func Test_search(t *testing.T) {
	fmt.Println(search([]int{1, 2, 3, 5, 9}, 9))
}
