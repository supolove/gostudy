package 链表

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/linked-list/196/doubly-linked-list/759/
*/

type DoublyNode struct {
	Val  int
	Next *DoublyNode
	Pre  *DoublyNode
}

type MyLinkedList struct {
	HeadNode *DoublyNode
	Len      int
}

/** Initialize your data structure here. */
func Constructor() *MyLinkedList {
	return &MyLinkedList{Len: 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.Len {
		return -1
	}
	cur := l.HeadNode
	var i = 0
	for {
		if i == index {
			return cur.Val
		}
		cur = cur.Next
		i++
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (l *MyLinkedList) AddAtHead(val int) {
	h := &DoublyNode{Val: val}
	if l.Len == 0 {
		l.HeadNode = h
		l.Len++
		return
	}
	h.Next = l.HeadNode
	l.HeadNode.Pre = h
	l.HeadNode = h
	l.Len++
}

/** Append a node of value val to the last element of the linked list. */
func (l *MyLinkedList) AddAtTail(val int) {
	t := &DoublyNode{Val: val}
	if l.Len == 0 {
		l.HeadNode = t
		l.Len++
		return
	}
	cur := l.HeadNode
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = t
	t.Pre = cur
	l.Len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (l *MyLinkedList) AddAtIndex(index int, val int) {

	if index == 0 {
		l.AddAtHead(val)
		return
	}
	if index == l.Len {
		l.AddAtTail(val)
		return
	}

	// 找到要添加的位置
	cur := l.HeadNode
	for i := 1; i < index; i++ {
		cur = cur.Next
	}

	// 先创建一个node
	t := &DoublyNode{Val: val}

	// 指定pre
	t.Pre = cur

	// 指定next
	t.Next = cur.Next

	// pre指向当前
	cur.Next = t

	// next指向当前的
	t.Next.Pre = t

	l.Len++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index >= l.Len {
		return
	}

	// 删除头
	if index == 0 {
		l.HeadNode = l.HeadNode.Next
		if l.HeadNode != nil {
			l.HeadNode.Pre = nil
		}
		l.Len--
		return
	}

	// 删除尾
	if index == l.Len-1 {
		cur := l.HeadNode
		for cur.Next != nil {
			cur = cur.Next
		}

		pre := cur.Pre
		pre.Next = nil
		l.Len--
		return
	}

	// 删除中间
	cur := l.HeadNode
	for i := 1; i <= index; i++ {
		cur = cur.Next
	}
	pre := cur.Pre
	next := cur.Next
	pre.Next = next
	next.Pre = pre
	l.Len--
}

func Test_MyLinkedList(t *testing.T) {
	obj := Constructor()

	//obj.AddAtHead(0)
	//obj.AddAtTail(1)
	//obj.AddAtTail(2)
	//obj.AddAtTail(4)
	//obj.AddAtIndex(3,3)
	//
	//obj.DeleteAtIndex(3)

	obj.AddAtHead(2)
	obj.DeleteAtIndex(1)
	obj.AddAtHead(2)
	obj.AddAtHead(7)
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(5)
	obj.AddAtTail(5)
	obj.Get(5)
	obj.DeleteAtIndex(6)
	obj.DeleteAtIndex(4)

	cur := obj.HeadNode
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println("")

}
