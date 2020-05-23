package 哈希表

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/205/practical-application-hash-map/817/
*/

// 超时
func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[i] = nums[i]
		for j := 1; j <= k; j++ {
			if i-j < 0 {
				break
			}
			if i-j >= 0 && m[i-j] == m[i] {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	m := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		ss, ok := m[nums[i]]
		if ok {
			for j := 0; j < len(ss); j++ {
				if i-ss[j] <= k {
					return true
				}
			}
		}

		if ok {
			s := m[nums[i]]
			s = append(s, i)
			m[nums[i]] = s
		} else {
			s := []int{i}
			m[nums[i]] = s
		}
	}
	return false
}

func Test_containsNearbyDuplicate(t *testing.T) {
	fmt.Println(containsNearbyDuplicate2([]int{1, 0, 1, 1}, 1))
}
