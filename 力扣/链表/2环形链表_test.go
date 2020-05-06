package 链表

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/linked-list/194/two-pointer-technique/744/
*/

func hasCycle(head *ListNode) bool {
	cur := head
	runer := head
	for cur != nil {

		runer = runer.Next
		if runer == cur {
			return true
		}
		if runer != nil {
			runer = runer.Next
		}
		if runer == cur {
			return true
		}

		if runer == nil {
			return false
		}
		cur = cur.Next
	}

	return false
}

func hasCycle2(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func Test_hasCycle(t *testing.T) {
	node := createSimpleListNode()
	fmt.Println(hasCycle2(node))
}
