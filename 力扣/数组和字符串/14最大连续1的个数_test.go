package 数组和字符串

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/201/two-pointer-technique/788/
*/

func findMaxConsecutiveOnes(nums []int) int {
	k := 0
	count := 0
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			if count < i-k && nums[k] != 0 {
				count = i - k
			}
			k = i
		} else {
			if nums[k] == 0 {
				k = i
			}
		}
	}

	if nums[i-1] == 0 {
		i--
	}

	if count <= i-k {
		count = i - k
	}

	return count
}

func Test_findMaxConsecutiveOnes(t *testing.T) {
	fmt.Println(findMaxConsecutiveOnes([]int{0}))
}
