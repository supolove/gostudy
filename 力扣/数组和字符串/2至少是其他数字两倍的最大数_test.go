package 数组和字符串

import (
	"fmt"
	"testing"
)

/**
https://leetcode-cn.com/explore/learn/card/array-and-string/198/introduction-to-array/771/
*/

func dominantIndex(nums []int) int {
	large := 0
	temp := 0
	idx := 0
	for midx, i := range nums {
		if large < i {
			temp = large
			large = i
			idx = midx
		} else if temp < i {
			temp = i
		}
	}

	if large >= temp*2 {
		return idx
	}

	return -1
}

func Test_dominantIndex(t *testing.T) {
	//nums := []int{3, 6, 1, 0}
	nums := []int{0, 0, 3, 2}
	fmt.Println(dominantIndex(nums))
}
