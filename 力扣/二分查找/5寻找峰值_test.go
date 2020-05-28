package 二分查找

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/binary-search/210/template-ii/841/
*/

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}

	if nums[0] > nums[1] {
		return 0
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	left, right := 0, len(nums)-1
	for left < right {
		if left == 0 {
			left++
			continue
		}
		if right == len(nums)-1 {
			right--
			continue
		}

		if nums[left-1] < nums[left] && nums[left] > nums[left+1] {
			return left
		}

		if nums[right-1] < nums[right] && nums[right] > nums[right+1] {
			return right
		}

		left++
		right--
	}
	if left == right && nums[left-1] < nums[left] && nums[left] > nums[left+1] {
		return left
	}

	return -1
}

func findPeakElement2(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if mid == l {
			if nums[l] < nums[r] {
				return r
			}
			return l
		}
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] > nums[mid-1] && nums[mid] < nums[mid+1] {
			l = mid
		} else if nums[mid] < nums[mid-1] && nums[mid] > nums[mid+1] {
			r = mid
		} else {
			r = mid // l = mid
		}
	}
	return 0
}

func Test_findPeakElement(t *testing.T) {
	fmt.Println(findPeakElement([]int{1, 2, 3}))
}
