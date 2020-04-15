package 数组和字符串

import (
	"fmt"
	"testing"
)

/**
  螺旋矩阵
https://leetcode-cn.com/explore/learn/card/array-and-string/199/introduction-to-2d-array/775/
*/

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}
	pValue := func(x, y int) int {
		return x*10000 + y
	}
	visited := map[int]bool{}
	isValid := func(x, y int) bool {
		_, ok := visited[pValue(x, y)]
		return !ok && 0 <= x && x < len(matrix) && 0 <= y && y < len(matrix[0])
	}
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	result := make([]int, 0)

	x := 0
	y := 0
	d := 0
	for {
		if isValid(x, y) {
			visited[pValue(x, y)] = true
			result = append(result, matrix[x][y])
		} else {
			x -= dir[d][0]
			y -= dir[d][1]

			d++
			if d == 4 {
				d = 0
			}
			if len(visited) == len(matrix)*len(matrix[0]) {
				break
			}
		}

		x += dir[d][0]
		y += dir[d][1]

	}
	return result
}

func Test_spiralOrder(t *testing.T) {
	value := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//fmt.Println(value[1][0])
	fmt.Println(spiralOrder(value))
}
