package 哈希表

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/hash-table/204/practical-application-hash-set/806/
*/

func singleNumber(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			delete(m, v)
			continue
		}
		m[v] = true
	}
	for k := range m {
		return k
	}
	return 0
}

// 精要算法
func singleNumber2(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func Test_sigleNumber(t *testing.T) {
	singleNumber2([]int{10, 10, 2, 3, 3})
}
