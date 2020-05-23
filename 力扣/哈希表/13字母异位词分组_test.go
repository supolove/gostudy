package 哈希表

import (
	"fmt"
	"sort"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/hash-table/206/practical-application-design-the-key/820/
*/
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, s := range strs {
		// 排序
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool {
			return b[i] > b[j]
		})
		// 添加或者创建数组
		arr, ok := m[string(b)]
		if ok {
			arr = append(arr, s)
			m[string(b)] = arr
		} else {
			arr := []string{s}
			m[string(b)] = arr
		}
	}
	var arrs [][]string
	for _, v := range m {
		arrs = append(arrs, v)
	}
	return arrs
}

func Test_groupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
