package 每日一题

import (
	"fmt"
	"strconv"
	"testing"
)

/*
https://leetcode-cn.com/problems/palindrome-number/
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	str := strconv.Itoa(x)
	l := len(str)
	for i := 0; i < l/2; i++ {
		left := str[i]
		right := str[l-i-1]
		if left != right {
			return false
		}
	}
	return true
}

func Test_isPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(-121))
}
