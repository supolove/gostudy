package _0_06

/*
https://leetcode-cn.com/problems/3sum/
*/

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	var ans [][]int
	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func Test_threeSum(t *testing.T) {
	//threeSum([]int{-2,0,1,1,2})
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, 8, -4}))
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}
