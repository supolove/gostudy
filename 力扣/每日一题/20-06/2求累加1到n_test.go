package _0_06

import (
	"fmt"
	"math"
	"testing"
)

/*
https://leetcode-cn.com/problems/qiu-12n-lcof/
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
*/

func sumNums(n int) int {
	return (int(math.Pow(float64(n), 2.0)) + n) >> 1
}

func sumNums2(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

func Test_sumNums(t *testing.T) {
	fmt.Println(12 >> 1)
	fmt.Println(sumNums(3))
}
