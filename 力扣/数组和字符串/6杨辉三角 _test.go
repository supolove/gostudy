package 数组和字符串

import "testing"

/*
  杨辉三角
https://leetcode-cn.com/explore/learn/card/array-and-string/199/introduction-to-2d-array/776/
*/

func generate(numRows int) [][]int {
	result := [][]int{}

	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		if i > 0 {
			up := result[i-1]
			for j := 0; j < len(up); j++ {
				arr[j] += up[j]
				arr[j+1] += up[j]
			}
			result = append(result, arr)
		} else {
			arr[0] = 1
			result = append(result, arr)
		}
	}
	return result
}

func Test_generate(t *testing.T) {
	generate(5)
}
