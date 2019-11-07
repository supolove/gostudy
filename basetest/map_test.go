package basetest

import (
	"fmt"
	"testing"
)

// 字典，映射， key-value， key唯一
func TestMap(t *testing.T) {
	var m1 map[int]string
	fmt.Println(m1 == nil)

	m2 := map[int]string{}
	m2[1] = "shengli"

	m3 := make(map[int]string)
	fmt.Println(len(m3))

	m4 := make(map[string]string, 10)
	fmt.Println(m4)

	m5 := map[int]string{0: "a", 1: "b"}
	fmt.Println(m5)
	m5[1] = "nil"

	// 遍历
	for _, v := range m5 {
		fmt.Println(v)
	}

	// 判断key是否存在
	if _, ok := m5[2]; ok {
	} else {

	}

	// 删除
	delete(m5, 1)

	// wcFunc() 单词统计函数

}
