package 数据结构

import (
	list2 "container/list"
	"fmt"
	"testing"
)

// 链表
// 链表就是一个有 prev 和 next 指针的数组了。它维护两个 type
func TestList(t *testing.T) {
	list := list2.New()
	list.PushBack(1)
	list.PushBack(2)

	fmt.Printf("len: %v\n", list.Len())
	fmt.Printf("first: %#v\n", list.Front())
	fmt.Printf("second: %#v\n", list.Front().Next())
}
