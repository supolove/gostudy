package 栈

import (
	"fmt"
	"testing"
)

/**
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

https://leetcode-cn.com/explore/learn/card/queue-stack/220/conclusion/892/
*/

type PositionUpdateMatrix struct {
	X     int
	Y     int
	index int
}

func updateMatrix(matrix [][]int) [][]int {
	queue := make([]*PositionUpdateMatrix, 0)

	key := func(x, y int) int {
		return x*100000 + y
	}

	isValid := func(x, y int, visited map[int]bool) bool {
		_, ok := visited[key(x, y)]
		return x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && !ok
	}

	// x是纵轴，y是横轴
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != 1 {
				continue
			}
			queue = append(queue, &PositionUpdateMatrix{i, j, 0})
			cur := queue[0]
			visited := map[int]bool{}
			for len(queue) > 0 {
				item := queue[0]
				queue = queue[1:]
				visited[key(item.X, item.Y)] = true

				if matrix[item.X][item.Y] == 0 {
					matrix[cur.X][cur.Y] = item.index
					queue = queue[0:0]
					break
				}

				if isValid(item.X+1, item.Y, visited) {
					queue = append(queue, &PositionUpdateMatrix{item.X + 1, item.Y, item.index + 1})
				}
				if isValid(item.X-1, item.Y, visited) {
					queue = append(queue, &PositionUpdateMatrix{item.X - 1, item.Y, item.index + 1})
				}
				if isValid(item.X, item.Y+1, visited) {
					queue = append(queue, &PositionUpdateMatrix{item.X, item.Y + 1, item.index + 1})
				}
				if isValid(item.X, item.Y-1, visited) {
					queue = append(queue, &PositionUpdateMatrix{item.X, item.Y - 1, item.index + 1})
				}

			}
		}
	}
	return matrix
}

func Test_updateMatrix(t *testing.T) {
	//matrix := [][]int{{0,0,0},{0,1,0}, {1,1,1}}
	matrix := [][]int{{0}}
	fmt.Println(updateMatrix(matrix))
}
