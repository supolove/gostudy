package 链表

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/linked-list/194/two-pointer-technique/747/
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || (n == 1 && head.Next == nil) {
		return nil
	}

	slow := head
	fast := head
	idx := 0
	for fast != nil {
		if idx >= n+1 {
			slow = slow.Next
			idx--
		}
		fast = fast.Next
		idx++
	}
	if slow == head && n == idx {
		return slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	listNode := &ListNode{Val: 0}
	listNode.Next = head
	fast, slow := listNode, listNode
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	return listNode.Next
}

func Test_removeNthFromEnd(t *testing.T) {
	n1 := createNode(1)
	n2 := createNode(2)
	//n3 := createNode(3)
	//n4 := createNode(4)
	//n5 := createNode(5)
	n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5

	printListNode(removeNthFromEnd(n1, 2))
}
