package 数组和字符串

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/198/introduction-to-array/1412/
*/

/*
使用二分查找会更快
*/

func searchInsert(nums []int, target int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v >= target {
			r = i
			break
		}

		if i == len(nums)-1 {
			r = i + 1
		}

		if v < target {
			continue
		}
	}
	return r
}

func Test_searchInsert(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}
