package 栈

import (
	"fmt"
	"testing"
)

/*
每日温度
根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。
如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，
你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

提示：该题目的意思是，当天的温度，需要等待多少天才比当前温度高。
比如，73度要再等一天就，所以是1

*/

// 我的做法
func dailyTemperatures(T []int) []int {
	if len(T) == 1 {
		return []int{0}
	}

	result := make([]int, len(T))

	idx := 0
	for len(T)-1 > 0 {
		v := T[0]
		T = T[1:]
		i := 0
		for _, t := range T {
			i++
			if t > v {
				break
			}
			if i == len(T) {
				i = 0
			}
		}
		result[idx] = i
		idx++
	}
	return result
}

// 栈做法
func dailyTemperatures1(T []int) []int {

	result := make([]int, len(T))
	stack := make([]int, 0)
	for i := 0; i < len(T); i++ {
		for len(stack) != 0 && T[stack[len(stack)-1]] < T[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[index] = i - index
		}
		stack = append(stack, i)
	}

	return result
}

func Test_dailyTemperatures(t *testing.T) {
	//a := []int{73, 72, 75, 71, 69, 72, 76, 73}
	a := []int{78, 75, 76, 79}
	fmt.Println(dailyTemperatures(a))
	fmt.Println(dailyTemperatures1(a))
}
