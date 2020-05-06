package 链表

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/linked-list/194/two-pointer-technique/745/
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	m := map[*ListNode]bool{}
	cur := head
	for cur != nil {
		m[cur] = true
		cur = cur.Next
		_, ok := m[cur]
		if ok {
			return cur
		}
	}
	return nil

	//runner := head
	//for runner != nil && runner.Next != nil  {
	//	cur = cur.Next
	//	runner = runner.Next.Next
	//	if runner == cur {
	//		return cur.Next
	//	}
	//}
	//return nil
}

func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for slow, fast := head, head; slow != nil && fast != nil; {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
		if fast != nil && fast == slow {
			slow = head
			for slow != fast {
				slow, fast = slow.Next, fast.Next
			}
			return fast
		}
	}
	return nil
}

func Test_detectCycle(t *testing.T) {
	node := createSimpleListNode()
	//node := &ListNode{Val:1}
	//node2 := &ListNode{Val:2}
	//node.Next = node2
	fmt.Println(detectCycle(node))
}
