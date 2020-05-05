package 数组和字符串

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/202/conclusion/796/
*/

func moveZeroes(nums []int) {

	s := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[s] = nums[i]
			s++
		}
	}
	for i := s; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}

func moveZeroes2(nums []int) {
	firstZeroIndex := -1
	zeroCount := 0
	for idx, num := range nums {
		//2.遇到0，将0与后面的值交换
		//2.1当有多个0时，将第一个0与非零值交换
		if num == 0 {
			zeroCount++
			if zeroCount == 1 { //找到第1个0
				firstZeroIndex = idx
			}
		} else if zeroCount > 0 {
			nums[firstZeroIndex], nums[idx] = nums[idx], nums[firstZeroIndex]
			firstZeroIndex++
		}
	}
	fmt.Println(nums)
}

func Test_moveZeroes(t *testing.T) {
	moveZeroes2([]int{0, 1, 0, 3, 12})
}
