package _0_06

import (
	"fmt"
	"strings"
	"testing"
)

/*
https://leetcode-cn.com/problems/valid-palindrome/
*/

func validPalindrome(s string) bool {
	s = strings.ToLower(s)
	length := len(s)

	isValid := func(b byte) bool {
		return ('a' <= b && b <= 'z') || ('0' <= b && b <= '9')
	}

	for i, j := 0, length-1; i < j; {
		if !isValid(s[i]) {
			i++
			continue
		}
		if !isValid(s[j]) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func Test_validPalindrome(t *testing.T) {
	fmt.Println(validPalindrome("A man, a plan, a canal: Panama"))
}
