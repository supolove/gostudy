package 哈希表

import (
	"fmt"
	"testing"
)

func firstUniqChar(s string) int {
	m := map[rune]int{}
	cache := map[rune]bool{}
	for i, n := range s {
		_, ok := cache[n]
		if ok {
			delete(m, n)
		} else {
			m[n] = i
			cache[n] = true
		}
	}

	min := -1
	for _, v := range m {
		if min == -1 {
			min = v
		} else if min > v {
			min = v
		}
	}
	return min
}

func Test_firstUniqChar(t *testing.T) {
	fmt.Println(firstUniqChar("loveleetcode"))
}
