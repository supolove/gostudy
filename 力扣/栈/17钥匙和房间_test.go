package 栈

import (
	"fmt"
	"testing"
)

/*

钥匙和房间
https://leetcode-cn.com/explore/learn/card/queue-stack/220/conclusion/893/
*/

// 题意理解错误
func canVisitAllRooms1(rooms [][]int) bool {
	m := map[int]bool{}

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			if i != rooms[i][j] {
				m[rooms[i][j]] = true
			}
		}
	}

	for i := 1; i < len(rooms); i++ {
		_, ok := m[i]
		if !ok {
			return false
		}
	}

	return true
}

// 队列实现方式
func canVisitAllRooms2(rooms [][]int) bool {
	visited := map[int]bool{}

	queue := make([]int, 0)
	for _, key := range rooms[0] {
		queue = append(queue, key)
	}
	visited[0] = true

	isValidRoom := func(key int) bool {
		_, ok := visited[key]
		return !ok && key < len(rooms)
	}

	for len(queue) > 0 {
		roomNumber := queue[0]
		queue = queue[1:]
		visited[roomNumber] = true

		for _, key := range rooms[roomNumber] {
			if isValidRoom(key) {
				queue = append(queue, key)
			}
		}
	}

	for i := 1; i < len(rooms); i++ {
		_, ok := visited[i]
		if !ok {
			return false
		}
	}

	return true
}

// 网友实现方式
var tmp map[int]int

func canVisitAllRooms(rooms [][]int) bool {
	tmp = make(map[int]int, 0)
	tmp[0] = 0
	add(rooms, 0)
	for i := 0; i < len(rooms); i++ {
		if _, ok := tmp[i]; !ok {
			return false
		}
	}
	return true
}
func add(rooms [][]int, k int) {
	for i := 0; i < len(rooms[k]); i++ {
		if _, ok := tmp[rooms[k][i]]; !ok {
			tmp[rooms[k][i]] = 0
			add(rooms, rooms[k][i])
		}

	}
}

func Test_canVisitAllRooms(t *testing.T) {
	// [[4],[3],[],[2,5,7],[1],[],[8,9],[],[],[6]]

	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	//rooms := [][]int{{1,2},{2},{3},{}}
	fmt.Println(canVisitAllRooms(rooms))
}
