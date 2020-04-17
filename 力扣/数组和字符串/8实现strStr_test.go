package 数组和字符串

import (
	"fmt"
	"strings"
	"testing"
)

/**
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。


https://leetcode-cn.com/submissions/detail/63766240/
*/

// 按照题意不应该这样简单的实现
func strStr1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	return strings.Index(haystack, needle)
}

// 实现方法
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for j := 0; j < len(haystack); j++ {
		if haystack[j] == needle[0] {
			r := 0
			for i := 0; i < len(needle); i++ {
				if j+i >= len(haystack) {
					return -1
				}
				if needle[i] != haystack[j+i] {
					break
				} else {
					r++
				}
			}
			if r == len(needle) {
				return j
			}
		}
	}

	return -1
}

func Test_strStr(t *testing.T) {
	fmt.Println(strStr("hello", "ll"))
}
