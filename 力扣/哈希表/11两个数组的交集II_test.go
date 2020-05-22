package 哈希表

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/205/practical-application-hash-map/816/
*/

func intersect(nums1 []int, nums2 []int) []int {
	len1 := len(nums1)
	len2 := len(nums2)
	var bl = make([]bool, len2)
	var al []int
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if nums1[i] == nums2[j] && bl[j] == false {
				al = append(al, nums1[i])
				bl[j] = true
				break
			}
		}
	}
	return al
}

func intersect2(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		if val, ok := m[nums1[i]]; !ok {
			m[nums1[i]] = 1
		} else {
			val++
			m[nums1[i]] = val
		}
	}
	var al []int
	val := 0
	for i := 0; i < len(nums2); i++ {
		val = m[nums2[i]]
		_, ok := m[nums2[i]]
		if ok && val > 0 {
			al = append(al, nums2[i])
			val--
			m[nums2[i]] = val
		}
	}
	return al
}

func Test_intersect(t *testing.T) {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	//fmt.Println(intersect([]int{1}, []int{1,1}))
	//fmt.Println(intersect2([]int{6,4,9,5}, []int{9,4,9,8,4}))
}
