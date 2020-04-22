package 数组和字符串

import (
	"fmt"
	"testing"
)

/**

数组拆分
https://leetcode-cn.com/explore/learn/card/array-and-string/201/two-pointer-technique/784/
给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。

一个10万字节的数组排序
冒泡排序使用600毫秒
快速排序4毫秒
*/

func arrayPairSum(nums []int) int {
	// 从小到大冒泡排序
	for i := len(nums); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if nums[j+1] < nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	result := 0
	// 从大到小冒泡排序
	for i := len(nums); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if nums[j+1] > nums[j] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if i%2 != 0 {
			result += nums[i]
		}
	}

	return result
}

// 使用快速排序
func arrayPairSum2(nums []int) int {

	partition := func(a []int, low, high int) int {
		x := a[low]
		i := low
		j := high

		for i < j {
			for i < j && a[j] >= x {
				j--
			}
			for i < j && a[i] <= x {
				i++
			}
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
		}
		a[i], a[low] = a[low], a[i]
		return i
	}

	var quicksort func(a []int, low, high int)
	quicksort = func(a []int, low, high int) {
		if low < high {
			i := partition(a, low, high)
			quicksort(a, low, i-1)
			quicksort(a, i+1, high)
		}
	}

	quicksort(nums, 0, len(nums)-1)

	result := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			result += nums[i]
		}
	}

	return result
}

func arrayPairSum3(nums []int) int {
	bucket, res, flag := make([]int, 20001), 0, false
	for i := range nums {
		bucket[nums[i]+10000]++
	}
	for i := range bucket {
		if bucket[i] == 0 {
			continue
		}
		if flag {
			bucket[i]--
			flag = false
		}
		if bucket[i]&1 == 1 {
			bucket[i]++
			flag = true
		}
		res += (i - 10000) * bucket[i] / 2
	}
	return res
}

func Test_arrayPairSum(t *testing.T) {
	fmt.Println(arrayPairSum2([]int{4, 3, 2, 1}))
}
