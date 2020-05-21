package 哈希表

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/hash-table/205/practical-application-hash-map/813/
*/

func findRestaurant(list1 []string, list2 []string) []string {
	m := map[string]int{}
	for i, s := range list1 {
		m[s] = i
	}

	idx := -1
	var result []string
	for i, ss := range list2 {
		v, ok := m[ss]
		if ok {
			if idx == -1 {
				idx = v + i
				result = append(result, ss)
			} else if idx == v+i {
				result = append(result, ss)
			} else if idx > v+i {
				idx = v + i
				result = []string{ss}
			}
		}
	}
	return result
}

func Test_findRestaurant(t *testing.T) {

}
