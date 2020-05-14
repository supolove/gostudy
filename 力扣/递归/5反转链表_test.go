package 递归

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/orignial/card/recursion-i/257/recurrence-relation/1210/
*/

func digui3(head, cur *ListNode) (*ListNode, *ListNode) {
	if cur.Next != nil {
		var mcur *ListNode
		head, mcur = digui3(head, cur.Next)
		mcur.Next = cur
	} else {
		head = cur
	}
	return head, cur
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	h, _ := digui3(head, cur)
	head.Next = nil
	return h
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func Test_reverseList(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	re := reverseList2(n1)
	for re != nil {
		fmt.Println(re.Val)
		re = re.Next
	}
}
