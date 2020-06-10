package 每日一题

import (
	"fmt"
	"strconv"
	"testing"
)

/*
https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
*/

func translateNum(num int) int {
	source := strconv.Itoa(num)
	large := num / 2
	sum := 1
	for i := 1; i <= large; i++ {
		for j := 0; j < len(source)-1; j++ {
			r := (source[j]-'0')*10 + source[j+1] - '0'
			if r < 26 {
				sum++
			}
		}
	}
	return sum
}

func translateNum2(num int) int {
	if num <= 9 {
		return 1
	}
	//获取输入数字的余数，然后递归的计算翻译方法
	ba := num % 100
	//如果小于等于9或者大于等于26的时候，余数不能按照2位数字组合，
	// 比如56，只能拆分为5和6；反例25，可以拆分为2和5，也可以作为25一个整体进行翻译。
	if ba <= 9 || ba >= 26 {
		return translateNum2(num / 10)
	} else {
		// ba=[10, 25]时，既可以当做一个字母，也可以当做两个字母
		return translateNum2(num/10) + translateNum2(num/100)
	}
}

// 动态规划，官方解法
func translateNum3(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}

func Test_translateNum(t *testing.T) {
	fmt.Println(translateNum3(12258))
}
