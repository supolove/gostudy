package 哈希表

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/hash-table/207/conclusion/825/
*/

func numJewelsInStones(J string, S string) int {
	m := map[rune]bool{}
	for _, j := range J {
		m[j] = true
	}

	c := 0
	for _, s := range S {
		if _, ok := m[s]; ok {
			c++
		}
	}
	return c

}

func Test_numJewelsInStones(t *testing.T) {

}
