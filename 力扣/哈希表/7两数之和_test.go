package 哈希表

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/hash-table/205/practical-application-hash-map/811/
*/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// target - nums[i] = nums[j]
func twoSum2(nums []int, target int) []int {
	set := make(map[int]int)
	for index, num := range nums {
		_, exist := set[target-num]
		if exist {
			return []int{index, set[target-num]}
		} else {
			set[num] = index
		}
	}

	return nil
}

func Test_twoSum(t *testing.T) {

}
