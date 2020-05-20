package 哈希表

/*
https://leetcode-cn.com/explore/learn/card/hash-table/204/practical-application-hash-set/807/
*/

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, n := range nums1 {
		m[n] = 1
	}
	var arr []int
	for _, n := range nums2 {
		v, ok := m[n]
		if ok && v == 1 {
			arr = append(arr, n)
			m[n] = m[n] + 1
		}
	}
	return arr
}
