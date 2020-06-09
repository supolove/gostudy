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
	ba := num % 100
	if ba <= 9 || ba >= 26 {
		return translateNum2(num / 10)
	} else {
		return translateNum2(num/10) + translateNum2(num/100)
	}
}

func Test_translateNum(t *testing.T) {
	fmt.Println('1' - '0')
	fmt.Println('2')
	fmt.Println('9')
	fmt.Println('a')
}
