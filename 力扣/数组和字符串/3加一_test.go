package 数组和字符串

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/198/introduction-to-array/772/
*/

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 != 10 {
			digits[i] = digits[i] + 1
			return digits
		}

		digits[i] = 0
		if i == 0 {
			r := make([]int, len(digits)+1)
			r[0] = 1
			return r
		}
	}

	return digits
}

func Test_plusOne(t *testing.T) {
	fmt.Println(plusOne([]int{9}))
}
