package 数组和字符串

import (
	"fmt"
	"testing"
)

/**
双指针-反转字符串
https://leetcode-cn.com/explore/learn/card/array-and-string/201/two-pointer-technique/783/
*/

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		t := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = t
	}
	fmt.Println(s)
}

func reverseString2(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Test_reverseString(t *testing.T) {
	fmt.Println([]byte{'h', 'e', 'l', 'l', '0'})
	reverseString([]byte{'h', 'e', 'l', 'l', '0'})
}
