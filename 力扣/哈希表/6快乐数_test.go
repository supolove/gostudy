package 哈希表

import (
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/204/practical-application-hash-set/808/
*/

func isHappy(n int) bool {

	m := map[int]bool{}
	caculate := func(value int) int {
		v := 10
		sum := 0
		for v <= value*10 {
			if value == 0 {
				continue
			}
			s := value % v
			vv := (s * 10 / v) * (s * 10 / v)
			sum += vv
			value = value - s
			v *= 10
		}
		return sum
	}
	for {
		v := caculate(n)
		if v == 1 {
			break
		}
		_, ok := m[v]
		if ok {
			return false
		}
		n = v
		m[v] = true
	}
	return true
}

func Test_isHappy(t *testing.T) {
	isHappy(801968324)
}
