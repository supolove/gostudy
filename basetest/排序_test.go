package basetest

import (
	"fmt"
	"sort"
	"testing"
)

type Student struct {
	Name  string
	Score int
}

// 数组类型
type Students []Student

/*
	继承sort的Interface
*/

func (s Students) Len() int {
	return len(s)
}

// 从小到大排序
func (s Students) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestSort(t *testing.T) {
	s := Students{
		Student{
			Name:  "小明",
			Score: 10,
		},
		Student{
			Name:  "小强",
			Score: 99,
		},
		Student{
			Name:  "苏大强",
			Score: 59,
		},
		Student{
			Name:  "张大山",
			Score: 80,
		},
	}
	// 执行排序
	sort.Sort(s)
	fmt.Println(s)
}
