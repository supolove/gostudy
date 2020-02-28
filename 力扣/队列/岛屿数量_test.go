package 队列

import (
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

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	x, y := 0, 0
	q := make(chan int, 20)
	m := map[int]int{}

	// 使用广度搜索
	go func() {
		for {

			if grid[x][y] == 1 {
				m[1000*x+y] = 1
				grid[x][y] = 0
			}

			if x < len(grid)-1 && grid[x+1][y] == 1 {
				x = x + 1
				continue
			}

			if y < len(grid[0])-1 && grid[x][y+1] == 1 {
				y = y + 1
				continue
			}

			if x == len(grid)-1 && y == len(grid[0])-1 {
				x = 0
				y = 0
				continue
			}

		}
	}()
	go func() {
		for {
			posiction := <-q
			x = posiction / 1000
			y = posiction % 1000
			println("x:", x, "y:", y)
		}
	}()

	return 0
}

func addToQueue(x, y int, m map[int]int, q chan int) {
	println("=== x:", x, "y:", y)
	if _, ok := m[x*1000+y]; !ok {
		q <- x*1000 + y
		m[x*1000+y] = 1
	}
}

func searchIsLands(x, y int, grid [][]byte) {

}

func TestDaoyu(t *testing.T) {
	b := [][]byte{
		{0, 0, 1},
		{1, 1, 0},
		{0, 1, 0},
	}
	numIslands(b)

	for {

	}
}
