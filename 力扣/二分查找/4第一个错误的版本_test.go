package 二分查找

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/binary-search/210/template-ii/839/
*/

func isBadVersion(version int) bool {
	return version >= 1
}

func firstBadVersion(n int) int {
	if n <= 0 {
		return -1
	}
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		midb := isBadVersion(mid)
		if midb && !isBadVersion(mid-1) {
			return mid
		} else if midb {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left == n && isBadVersion(left) {
		return left
	}
	return -1
}

func Test_firstBadVersion(t *testing.T) {
	fmt.Println(firstBadVersion(1))
}
