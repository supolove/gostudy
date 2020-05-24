package 哈希表

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/207/conclusion/826/
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := map[uint8]int{}
	cache := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		v, ok := cache[s[i]]
		if !ok {
			cache[s[i]] = i
		} else {
			if len(cache) > len(m) {
				m = cache
			}
			cache = map[uint8]int{}
			i = v
		}
	}

	if len(cache) > len(m) {
		return len(cache)
	}

	return len(m)
}

func lengthOfLongestSubstring2(s string) int {
	n := len(s)
	res := 0
	i := 0
	index := [128]int{}

	// 计算两个相同元素之间的距离
	for j := 0; j < n; j++ {
		i = max(index[s[j]], i)
		res = max(res, j-i+1)
		index[s[j]] = j + 1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring2("abccabcbb"))
}
