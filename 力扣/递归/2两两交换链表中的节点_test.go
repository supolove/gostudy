package 递归

import (
	"fmt"
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(head, pre, cur *ListNode) *ListNode {
	if cur == nil || cur.Next == nil {
		return head
	}

	t := cur.Next.Next
	next := cur.Next
	next.Next = cur
	cur.Next = t
	if pre != nil {
		pre.Next = next
	} else {
		head = next
	}
	return swap(head, cur, cur.Next)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	head = swap(head, nil, cur)
	return head
}

func Test_swapPairs(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n := swapPairs(n1)
	for n != nil {
		fmt.Print(n.Val, " ")
		n = n.Next
	}
}
