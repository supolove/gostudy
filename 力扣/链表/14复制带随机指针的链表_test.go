package 链表

import "testing"
import . "gostudy/力扣/链表/node14"

/*
https://leetcode-cn.com/explore/learn/card/linked-list/197/conclusion/766/
*/

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	m := map[*Node]*Node{}
	cur := head
	newHead := &Node{Val: cur.Val}
	newCur := newHead
	m[cur] = newHead
	//cur = cur.Next
	//m[cur.Val] = newHead
	for cur != nil {
		// 先创建node
		newNode, ok := m[cur]
		if !ok {
			newNode = &Node{Val: cur.Val}
			m[cur] = newNode
		}

		// 判断ｒａｎｄｏｍ
		if cur.Random != nil {
			n, ok := m[cur.Random]
			if !ok {
				r := &Node{Val: cur.Random.Val}
				newNode.Random = r
				m[cur.Random] = r
			} else {
				newNode.Random = n
			}
		}

		if newNode != newHead {
			newCur.Next = newNode
			newCur = newCur.Next
		}
		cur = cur.Next
	}

	return newHead
}

func Test_copyRandomList(t *testing.T) {
	//n1 := &Node{Val:7}
	//n2 := &Node{Val:13}
	//n3 := &Node{Val:11}
	//n4 := &Node{Val:10}
	//n5 := &Node{Val:1}
	//n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5
	//
	//n2.Random = n1
	//n3.Random = n5
	//n4.Random = n3
	//n5.Random = n1
	//n1.Random = nil

	n1 := &Node{Val: 3}
	n2 := &Node{Val: 3}
	n3 := &Node{Val: 3}
	n1.Next = n2
	n2.Next = n3

	n2.Random = n1

	copyRandomList(n1)
}
