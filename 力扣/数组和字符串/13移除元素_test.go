package 数组和字符串

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/201/two-pointer-technique/787/
*/

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[k] = nums[i]
			k++
		}
	}
	return k + 1
}

func Test_removeElement(t *testing.T) {

}
