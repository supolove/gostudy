package 链表

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/linked-list/195/classic-problems/754/
*/

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	last := head
	idx := 1
	for last.Next != nil {
		last = last.Next
		idx++
	}

	s := idx - 1
	for s > 0 {
		cur := head
		for i := 0; i < s; i++ {
			cur = cur.Next
		}
		if head.Val != cur.Val {
			return false
		}
		head = head.Next
		s -= 2
	}
	return true
}

func Test_isPalindrome(t *testing.T) {
	node := createSimpleListNode()
	println(isPalindrome(node))
}
