package 每日一题

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/problems/daily-temperatures/
*/

func dailyTemperatures(T []int) []int {
	if T == nil {
		return []int{}
	}
	length := len(T)
	maxIdx := 0
	res := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		// 记录最大值，缩小搜索范围
		if maxIdx == 0 {
			res[i] = 0
			maxIdx = i
			continue
		}

		// 如果温度比最大的还大，那就等于0
		if T[i] >= T[maxIdx] {
			res[i] = 0
			maxIdx = i
		} else {
			// 在最大的范围查找比他大的温度
			for j := i; j < maxIdx+1; j++ {
				if T[i] < T[j] {
					res[i] = j - i
					break
				}
			}
		}
	}
	return res
}

// 递减栈
func dailyTemperatures2(T []int) []int {
	lens, top := len(T), 0
	ret := make([]int, lens)
	stack := make([]int, lens)
	stack[0] = lens - 1
	for i := lens - 2; i >= 0; i-- {
		for ; top >= 0 && T[stack[top]] <= T[i]; top-- {
		}
		if top >= 0 { //i 的右边存在比 T[i] 大的元素
			ret[i] = stack[top] - i
		}
		top++
		stack[top] = i
	}
	return ret
}

func Test_dailyTemperatures(t *testing.T) {
	fmt.Println(dailyTemperatures2([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
