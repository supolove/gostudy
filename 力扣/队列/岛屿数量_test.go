package 队列

import (
	"fmt"
	"testing"
)

/**
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。
一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。
你可以假设网格的四个边均被水包围。

示例1：
	输入:
	11110
	11010
	11000
	00000

	输出: 1

示例1：
	输入:
	11000
	11000
	00100
	00011

	输出: 3

思路：
只要找到1就必定有岛屿
找到1向右搜索，如果遇到1则合并再继续搜索，遇到0则结束
向下搜索，如果遇到1则合并再继续搜索，遇到0则结束
*/

type Position struct {
	X int
	Y int
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var q []Position

	isIsland := func(x, y int) bool {
		return 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) && grid[x][y] == '1'
	}

	// 将陆地旁边的陆地放入探险队列
	checkAndEnqueue := func(x, y int) {
		if isIsland(x-1, y) {
			q = append(q, Position{x - 1, y})
		}
		if isIsland(x+1, y) {
			q = append(q, Position{x + 1, y})
		}
		if isIsland(x, y-1) {
			q = append(q, Position{x, y - 1})
		}
		if isIsland(x, y+1) {
			q = append(q, Position{x, y + 1})
		}
	}

	ans := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 第一次碰到陆地
			if grid[i][j] == 1 {
				q = append(q, Position{i, j})
				ans++

				// 探索这块岛屿所有相连的陆地
				for len(q) != 0 {
					root := q[0]
					q = q[1:]
					if grid[root.X][root.Y] == 0 {
						continue
					}
					// 标记为已探索
					grid[root.X][root.Y] = 0
					// 将相连的陆地放入探险队列
					checkAndEnqueue(root.X, root.Y)
				}
			}
		}
	}
	return ans
}

func TestDaoyu(t *testing.T) {
	b := [][]byte{
		{0, 0, 1},
		{1, 1, 0},
		{0, 1, 0},
	}
	n := numIslands(b)
	fmt.Println(n)
}
