package 数据结构

import (
	list2 "container/list"
	"fmt"
	"testing"
)

// 双向连表链表
// 链表就是一个有 prev 和 next 指针的数组了。它维护两个 type
func TestList(t *testing.T) {
	list := list2.New()
	list.PushBack(1)
	list.PushBack(2)
	e := &list2.Element{}
	e.Value = 2
	list.MoveToFront(e)

	fmt.Printf("len: %v\n", list.Len())
	fmt.Printf("first: %#v\n", list.Front())
	fmt.Printf("second: %#v\n", list.Front().Next())
}
