package 栈

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

url:
https://leetcode-cn.com/explore/learn/card/queue-stack/219/stack-and-dfs/883/
*/

type Position struct {
	X int
	Y int
}

// 自定义栈实现
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	stack := make([]Position, 0)

	isIslandd := func(p Position, grid [][]byte) bool {
		return 0 <= p.X && p.X < len(grid) && 0 <= p.Y && p.Y < len(grid[0]) && grid[p.X][p.Y] == '1'
	}

	popStack := func() Position {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return v
	}

	pushStack := func(v Position) {
		stack = append(stack, v)
	}

	ans := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 第一次碰到陆地
			if grid[i][j] == '1' {
				ans++
				pushStack(Position{i, j})
				for len(stack) > 0 {
					p := popStack()
					grid[p.X][p.Y] = '0'
					if isIslandd(Position{p.X + 1, p.Y}, grid) {
						pushStack(Position{p.X + 1, p.Y})
					}
					if isIslandd(Position{p.X - 1, p.Y}, grid) {
						pushStack(Position{p.X - 1, p.Y})
					}
					if isIslandd(Position{p.X, p.Y + 1}, grid) {
						pushStack(Position{p.X, p.Y + 1})
					}
					if isIslandd(Position{p.X, p.Y - 1}, grid) {
						pushStack(Position{p.X, p.Y - 1})
					}
				}
			}
		}
	}
	return ans
}

func TestDaoyu(t *testing.T) {
	//b := [][]byte{
	//	{0, 0, 1},
	//	{1, 1, 0},
	//	{0, 1, 0},
	//}

	b := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'0', '1', '0', '1', '0'},
		{'1', '1', '1', '0', '0'},
		{'0', '1', '0', '0', '1'},
	}

	n := numIslands(b)
	fmt.Println(n)
}
