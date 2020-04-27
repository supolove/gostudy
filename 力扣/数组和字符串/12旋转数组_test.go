package 数组和字符串

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/202/conclusion/791/
*/

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	n := make([]int, 0)
	k = k % len(nums)
	for i := 0; i < len(nums); i++ {
		n = append(n, nums[i])
	}
	o := n[k-1:]
	p := n[:k-1]

	fmt.Println(o)
	fmt.Println(p)
}

func Test_rotate(t *testing.T) {
	rotate([]int{1, 2, 3, 4, 5, 6}, 2)
}
