package 哈希表

/*
https://leetcode-cn.com/explore/learn/card/hash-table/204/practical-application-hash-set/805/
*/

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
