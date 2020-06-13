package _0_06

import (
	"fmt"
	"strconv"
	"testing"
)

/*
https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
*/

/*
初始化方向向右
记录访问元素
如果向前和向下不能再走就结束
*/

// 方向定义
const (
	right = iota
	down
	left
	up
)

// key生成
func getKey(x, y int) string {
	xs := strconv.Itoa(x)
	ys := strconv.Itoa(y)
	return xs + "," + ys
}

// 行走
func walk(x, y, direction int) (int, int) {
	switch direction {
	case right:
		y++
	case down:
		x++
	case left:
		y--
	case up:
		x--
	}
	return x, y
}

// 变动方向
func changeDirection(direction int) int {
	if direction == up {
		return right
	}
	direction++
	return direction
}

// 判断下一步是否可以走
func canWalk(x, y, direction int, arr [][]int, m map[string]bool) bool {
	// 尝试往前行走
	xx, yy := walk(x, y, direction)
	// 判断是否走过
	_, ok := m[getKey(xx, yy)]
	if ok {
		return false
	}
	// 判断是否在范围之内
	if 0 <= xx && xx < len(arr) && 0 <= yy && yy < len(arr[0]) {
		return true
	}
	return false
}

func spiralOrder(matrix [][]int) []int {
	var res []int
	if matrix == nil || len(matrix) == 0 {
		return res
	}
	direction := right
	x := 0
	y := 0
	res = append(res, matrix[x][y])
	m := map[string]bool{}
	m[getKey(x, y)] = true
	for {
		// 判断前面是否可走
		if !canWalk(x, y, direction, matrix, m) {
			direction = changeDirection(direction)
			// 修改方向后是否可走，如果不行就结束了
			if !canWalk(x, y, direction, matrix, m) {
				break
			}
		}
		// 上面没问题，就往前走一步
		x, y = walk(x, y, direction)
		// 记录走过的
		m[getKey(x, y)] = true
		// 添加路径
		res = append(res, matrix[x][y])
	}
	return res
}

func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		directionIndex = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order
}

func Test_spiralOrder(t *testing.T) {
	r := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(r))
}
