package 数组和字符串

import (
	"fmt"
	"testing"
)

/**
对角线遍历
给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
https://leetcode-cn.com/explore/learn/card/array-and-string/199/introduction-to-2d-array/774/
*/

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	r := 0
	c := 0
	k := 0
	res := make([]int, 0)

	dirs := [][]int{{-1, 1}, {1, -1}}
	for i := 0; i < m*n; i++ {
		res = append(res, matrix[r][c])
		r += dirs[k][0]
		c += dirs[k][1]
		if r >= m {
			r = m - 1
			c += 2
			k = 1 - k
		}
		if c >= n {
			c = n - 1
			r += 2
			k = 1 - k
		}
		if r < 0 {
			r = 0
			k = 1 - k
		}
		if c < 0 {
			c = 0
			k = 1 - k
		}
	}
	return res
}

func Test_findDiagonalOrder(t *testing.T) {
	r := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(findDiagonalOrder(r))
}
