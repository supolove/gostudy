package 栈

import (
	"fmt"
	"testing"
)

/**
图像渲染
https://leetcode-cn.com/explore/learn/card/queue-stack/220/conclusion/891/
*/

type PositionFloodFill struct {
	X int
	Y int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	oldColor := image[sr][sc]
	visited := map[PositionFloodFill]bool{}

	// 构建一个队列，做BFS搜索
	queue := make([]PositionFloodFill, 0)

	curp := PositionFloodFill{sr, sc}
	// 判断是否像素点
	isXp := func(p PositionFloodFill) bool {
		_, ok := visited[p]
		return 0 <= p.X && p.X < len(image) && 0 <= p.Y && p.Y < len(image[0]) && image[p.X][p.Y] == oldColor && !ok
	}

	// 添加一个元素
	queue = append(queue, curp)
	visited[curp] = true

	// 位置行走
	for len(queue) > 0 {
		p := queue[0]
		image[p.X][p.Y] = newColor
		queue = queue[1:]
		visited[p] = true
		if isXp(PositionFloodFill{p.X + 1, p.Y}) {
			queue = append(queue, PositionFloodFill{p.X + 1, p.Y})
		}

		if isXp(PositionFloodFill{p.X - 1, p.Y}) {
			queue = append(queue, PositionFloodFill{p.X - 1, p.Y})
		}

		if isXp(PositionFloodFill{p.X, p.Y + 1}) {
			queue = append(queue, PositionFloodFill{p.X, p.Y + 1})
		}

		if isXp(PositionFloodFill{p.X, p.Y - 1}) {
			queue = append(queue, PositionFloodFill{p.X, p.Y - 1})
		}
	}

	return image
}

func Test_floodFill(t *testing.T) {
	//image := [][]int{{1,1,1},{1,1,0},{1,0,1}}
	image := [][]int{{0, 0, 0}, {0, 1, 1}}
	r := floodFill(image, 1, 1, 2)
	fmt.Println(r)
}
