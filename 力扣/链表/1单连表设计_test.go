package 链表

import (
	"fmt"
	. "gostudy/力扣/链表/node1"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/linked-list/193/singly-linked-list/741/
*/

type MySingleLinkedList struct {
	HeadNode *Node
	Len      int
}

/** Initialize your data structure here. */
func SConstructor() *MySingleLinkedList {
	return &MySingleLinkedList{Len: 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MySingleLinkedList) Get(index int) int {
	cur := this.HeadNode
	var i = 0
	for {
		if cur == nil {
			return -1
		}
		if i == index {
			return cur.Val
		}
		cur = cur.Next
		i++
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MySingleLinkedList) AddAtHead(val int) {
	h := &Node{Val: val}
	h.Next = this.HeadNode
	this.HeadNode = h
	this.Len++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MySingleLinkedList) AddAtTail(val int) {
	t := &Node{Val: val}
	cur := this.HeadNode
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = t
	this.Len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MySingleLinkedList) AddAtIndex(index int, val int) {

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	cur := this.HeadNode
	i := 0
	for cur != nil {
		i++
		if i == index && cur.Next != nil {
			t := &Node{Val: val}
			t.Next = cur.Next
			cur.Next = t
			this.Len++
			break
		} else if cur.Next == nil {
			cur.Next = &Node{Val: val}
			this.Len++
			break
		}
		cur = cur.Next
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MySingleLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.HeadNode = this.HeadNode.Next
		return
	}
	cur := this.HeadNode
	i := 0
	for cur != nil {
		i++
		if i == index && cur.Next != nil {
			cur.Next = cur.Next.Next
			this.Len--
		}
		cur = cur.Next
	}
}

func Test_singleLinkedList(t *testing.T) {
	fmt.Println("-----")
	// 0 1 2 3 4

	obj := SConstructor()
	/*
		obj.Val = 0
		cur := obj
		for i := 1; i < 10; i++ {
			n := &MyLinkedList{Val: i}
			cur.Next = n
			cur = n
		}
	*/
	//param_1 := obj.Get(2)
	//fmt.Println(param_1)

	//obj.AddAtHead(11)
	//fmt.Println(obj.Get(0))
	//obj.AddAtTail(11)
	//obj.AddAtIndex(2,11)
	//obj.DeleteAtIndex(2)

	//obj.AddAtHead(7)
	//obj.AddAtHead(2)
	//obj.AddAtHead(1)
	//obj.AddAtIndex(3,0)
	//obj.DeleteAtIndex(2)
	//obj.AddAtHead(6)
	//obj.AddAtTail(4)
	//fmt.Println(obj.Get(4))

	//obj.AddAtHead(1)
	//obj.AddAtTail(3)
	//obj.AddAtIndex(1,2)
	//obj.Get(1)
	//obj.DeleteAtIndex(0)
	//obj.Get(0)

	//obj.AddAtHead(2)
	//obj.DeleteAtIndex(1)
	//obj.AddAtHead(2)
	//obj.AddAtHead(7)
	//obj.AddAtHead(3)
	//obj.AddAtHead(2)
	//obj.AddAtHead(5)
	//obj.AddAtTail(5)
	//obj.DeleteAtIndex(6)
	//obj.DeleteAtIndex(4)

	// 38 66 61 76 26 37 8
	obj.AddAtHead(38)
	obj.AddAtTail(66)
	obj.AddAtTail(61)
	obj.AddAtTail(76)
	obj.AddAtTail(26)
	obj.AddAtTail(37)
	obj.AddAtTail(8)
	obj.AddAtIndex(5, 10)
	//obj.AddAtIndex(7,10)

	mc := obj.HeadNode
	for mc != nil {
		fmt.Print(mc.Val, " ")
		mc = mc.Next
	}
}
