package 链表

import (
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/linked-list/195/classic-problems/750/
*/

// 迭代方法
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		t := cur.Next
		cur.Next = cur.Next.Next
		t.Next = head
		head = t
	}

	return head
}

// 递归方法
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// newHead得到最后一个节点
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func Test_reverseList(t *testing.T) {
	node := createSimpleListNode()
	r := reverseList2(node)
	printListNode(r)
}
