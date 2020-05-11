package 链表

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/linked-list/197/conclusion/767/
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	cur := head
	idx := 1
	for cur.Next != nil {
		cur = cur.Next
		idx++
	}
	cur.Next = head

	k = k % idx
	tail := cur
	for i := 0; i < idx-k; k++ {
		tail = head
		head = head.Next
	}
	tail.Next = nil
	return head
}

func Test_rotateRight(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	rotateRight(n1, 2)
}
