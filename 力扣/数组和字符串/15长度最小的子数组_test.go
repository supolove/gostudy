package 数组和字符串

import (
	"fmt"
	"testing"
)

/**
https://leetcode-cn.com/explore/learn/card/array-and-string/201/two-pointer-technique/789/
*/

func minSubArrayLen(s int, nums []int) int {
	arr := make([]int, 0)
	max := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			max += nums[j]
			if max >= s {
				arr = append(arr, j-i+1)
				max = 0
				break
			}
		}
		if max != 0 {
			break
		}
	}
	l := len(nums)
	if len(arr) == 0 {
		return 0
	}
	for _, i := range arr {
		if l > i {
			l = i
		}
	}
	return l
}

func Test_minSubArrayLen(t *testing.T) {
	//213
	//[12,28,83,4,25,26,25,2,25,25,25,12]
	fmt.Println(minSubArrayLen(213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}))
	//fmt.Println(minSubArrayLen(7, []int{2,3,1,2,4,3}))

}
