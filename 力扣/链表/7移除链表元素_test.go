package 链表

import "testing"

/**
https://leetcode-cn.com/explore/learn/card/linked-list/195/classic-problems/752/
*/

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil || (head.Next == nil && val == head.Val) {
		return nil
	}

	pre := head
	cur := head
	for cur != nil {
		if cur.Val == val {
			if cur == head {
				head = cur.Next
				pre = head
				cur = head
				continue
			}
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}

	return head
}

func Test_removeElements(t *testing.T) {
	n := createSimpleListNode()

	printListNode(removeElements(n, 2))
}
