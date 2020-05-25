package 哈希表

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/207/conclusion/829/
*/

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var res []int
	for i := 0; i < k; i++ {
		curv := 0
		curk := 0
		for kk, v := range m {
			if v > curv {
				curk, curv = kk, v
			}
		}
		if curv != 0 {
			res = append(res, curk)
			delete(m, curk)
		}
	}

	return res
}

// 进行排序
func topKFrequent2(nums []int, k int) []int {
	numMap := make(map[int]int, len(nums))
	for i := range nums {
		numMap[nums[i]]++
	}
	type numsk struct {
		num   int
		count int
	}
	res := make([]int, 0, k)
	sk := make([]numsk, 0, len(numMap))
	for num, count := range numMap {
		sk = append(sk, numsk{num: num, count: count})
	}

	sort.Slice(sk, func(i, j int) bool {
		return sk[i].count > sk[j].count
	})
	for i := range sk {
		res = append(res, sk[i].num)
	}

	return res[:k]
}

func Test_topKFrequent(t *testing.T) {
	//fmt.Println(topKFrequent([]int{1,2},2))
	fmt.Println(topKFrequent([]int{1, 2}, 1))
}
