package 每日一题

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/problems/product-of-array-except-self/
*/

// 平民算法
func productExceptSelf(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		multiply := 1
		for j := len(nums) - 1; j >= 0; j-- {
			if i != j {
				multiply *= nums[j]
			}
		}
		res = append(res, multiply)
	}
	return res
}

// 左右乘积列表   1 2 3 4
func productExceptSelf2(nums []int) []int {
	length := len(nums)
	L := make([]int, len(nums))
	R := make([]int, len(nums))

	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	// 0:1 1:1  2:2 3:6
	// 4:4 3:12 2:24

	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	var answer []int
	for i := 0; i < length; i++ {
		answer = append(answer, L[i]*R[i])
	}

	return answer
}

func Test_productExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf2([]int{1, 2, 3, 4}))
}
