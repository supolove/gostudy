package 数组和字符串

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/array-and-string/198/introduction-to-array/1413/
*/

func mmin(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}

func mmax(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
func isOverlapping(arr1, arr2 []int) (bool, []int) {
	if (arr2[0] <= arr1[0] && arr1[0] <= arr2[1]) || (arr1[0] <= arr2[0] && arr2[0] <= arr1[1]) {
		return true, []int{mmin(arr1[0], arr2[0]), mmax(arr1[1], arr2[1])}
	}
	return false, nil
}

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return nil
	}
	var res [][]int
	for i := 0; i < len(intervals); i++ {
		isMerge := false
		for j := i + 1; j < len(intervals); j++ {
			b, arr := isOverlapping(intervals[i], intervals[j])
			if b {
				intervals[j] = arr
				isMerge = true
				break
			}
		}
		if isMerge == false {
			res = append(res, intervals[i])
		}
	}
	return res
}

func Test_merge(t *testing.T) {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {1, 10}}))
}
