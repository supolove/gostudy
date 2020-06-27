package 数组和字符串

import (
	"fmt"
	"testing"
)

/**
https://leetcode-cn.com/explore/learn/card/array-and-string/198/introduction-to-array/770/
*/

/*
集合：由一个或多个确定的元素所构成的整体。
特性：
1.集合里的元素类型不一定相同
2.集合里的元素没有顺序

列表：是一种数据项构成的有限序列，即按照一定的线性顺序，排列而成的数据项的集合

数组：
特征：
1.数组是列表的实现方式，它具有列表的特征
2.数组会用一些名为 索引 的数字来标识每项数据在数组中的位置
3.数组中的元素在内存中是连续存储的，且每个元素占用相同大小的内存


*/

// 只能算正数
func pivotIndex1(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	leftIdx := 0
	rightIdx := len(nums) - 1
	left := nums[leftIdx]
	right := nums[len(nums)-1]

	for rightIdx > leftIdx {
		if left > right {
			rightIdx--
			right += nums[rightIdx]
		} else {
			leftIdx++
			left += nums[leftIdx]
		}
		if left == right && rightIdx-leftIdx == 2 {
			return leftIdx + 1
		}

		if left == right && rightIdx-leftIdx == 1 {
			return leftIdx
		}
	}

	return -1
}

// 可以算负数
func pivotIndex(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	right := 0
	for _, i := range nums {
		right += i
	}

	// 吐数
	idx := 0
	left := 0
	for idx != len(nums)-1 {
		left += nums[idx]
		right -= nums[idx]

		if idx == 0 && right == 0 {
			return idx
		}

		if left == right-nums[idx+1] {
			return idx + 1
		}

		idx++
	}
	return -1
}

// 网友精要算法
func pivotIndex2(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	tmp := 0
	for i := 0; i < len(nums); i++ {
		// 去掉中间数是左边的两倍
		if tmp*2 == sum-nums[i] {
			return i
		}
		tmp += nums[i]
	}
	return -1
}

func Test_pivotIndex(t *testing.T) {

	// [-1,-1,-1,-1,-1,0]

	//nums := []int{1,7,3,6,5,6}
	//nums := []int{-1,-1,-1,-1,-1,0}
	//nums := []int{1,2,3,4,5}
	//nums := []int{-1,-1,1,1,0,0}
	nums := []int{-1, 1, -1, -1, 0, 1}
	fmt.Println(pivotIndex(nums))
}
