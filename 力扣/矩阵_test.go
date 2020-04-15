package 力扣

import (
	"fmt"
	"testing"
)

func Test_juzheng(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}

	fmt.Println(matrix[0][0])
	fmt.Println(matrix[0][1])
	// 先遍历x，再遍历y
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			fmt.Println(matrix[y][x])
		}
	}
}
