package 链表

import (
	"fmt"
	. "gostudy/力扣/链表/node13"
	"testing"
)

/**
https://leetcode-cn.com/explore/learn/card/linked-list/197/conclusion/764/
*/

func flatten(root *Node) *Node {
	cur := root
	for cur != nil {
		if cur.Child == nil {
			cur = cur.Next
			continue
		}
		next := cur.Next
		c := cur.Child
		cur.Child = nil
		n := flatten(c)
		cur.Next = n
		n.Prev = cur

		// 找到最后一个child节点
		for n.Next != nil {
			if n.Child != nil {
				c2 := n.Child
				n.Child = nil
				flatten(c2)
			}
			n = n.Next
		}

		if next == nil {
			break
		}
		n.Next = next
		next.Prev = n
		cur = next
	}
	return root
}

func Test_flatten(t *testing.T) {
	//n1 := &Node{Val:1}
	//n2 := &Node{Val:2}
	//n3 := &Node{Val:3}
	//n4 := &Node{Val:4}
	//n5 := &Node{Val:5}
	//n6 := &Node{Val:6}
	//
	//c1 := &Node{Val:7}
	//c2 := &Node{Val:8}
	//c3 := &Node{Val:9}
	//c4 := &Node{Val:10}
	//
	//cc1 := &Node{Val:11}
	//cc2 := &Node{Val:12}
	//
	//
	//n1.Next = n2; n2.Prev = n1
	//n2.Next = n3; n3.Prev = n2
	//n3.Next = n4; n4.Prev = n3
	//n4.Next = n5; n5.Prev = n4
	//n5.Next = n6; n6.Prev = n5
	//
	//c1.Next = c2; c2.Prev = c1
	//c2.Next = c3; c3.Prev = c2
	//c3.Next = c4; c4.Prev = c3
	//
	//cc1.Next = cc2; cc2.Prev = cc1
	//
	//n3.Child = c1; c2.Child = cc1

	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n1.Child = n2

	n := flatten(n1)
	for i := 0; n != nil; i++ {
		fmt.Println(n.Val, " ")
		n = n.Next
		if i > 100 {
			break
		}
	}
	fmt.Println(n)
}
