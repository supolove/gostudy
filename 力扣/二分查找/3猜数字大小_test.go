package 二分查找

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/binary-search/209/template-i/837/a
*/

func guess(num int) int {
	if num > 6 {
		return 1
	} else if num == 6 {
		return 0
	}
	return -1
}

func guessNumber(n int) int {
	left, right := 0, n
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(mid)
		r := guess(mid)
		if r == 0 {
			return mid
		} else if r == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func Test_guessNumber(t *testing.T) {
	fmt.Println(guessNumber(10))
}
