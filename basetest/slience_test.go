package basetest

import (
	"fmt"
	"testing"
)

func TestSlience(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := data[9:]
	// 创建长度容量为2的切片
	s2 := data[:2]

	fmt.Println(s2)
	// s1 会替换 s2的元素
	copy(s2, s1)

	fmt.Println(s2)
}
