package 数组和字符串

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/201/two-pointer-technique/785/
*/

// 最简单的做法
func twoSum1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := len(numbers) - 1; j > i; j-- {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for {
		sums := numbers[i] + numbers[j]
		if sums > target {
			j--
		} else if sums < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}
}

func Test_twoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
