package _0_06

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode-cn.com/problems/3sum-closest/
*/

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	closestNum := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			threeSum := nums[l] + nums[r] + nums[i]
			if absv(threeSum-target) < absv(closestNum-target) {
				closestNum = threeSum
			}
			if threeSum > target {
				r--
			} else if threeSum < target {
				l++
			} else {
				return target
			}

		}
	}
	return closestNum
}

func absv(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func Test_threeSumClosest(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 2))
}
