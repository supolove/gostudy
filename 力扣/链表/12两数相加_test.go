package 链表

import "testing"

/*
https://leetcode-cn.com/explore/learn/card/linked-list/197/conclusion/763/
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	pre := temp

	for l1 != nil && l2 != nil {
		v := 0
		if pre.Next != nil {
			v = l1.Val + l2.Val + 1
		} else {
			v = l1.Val + l2.Val
		}
		if v >= 10 {
			pre.Next = &ListNode{Val: v - 10}
			pre = pre.Next
			pre.Next = &ListNode{Val: 1}
		} else {
			pre.Next = &ListNode{Val: v}
			pre = pre.Next
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if pre.Next != nil {
			v := l1.Val + 1
			if v >= 10 {
				pre.Next = &ListNode{Val: v - 10}
				pre = pre.Next
				pre.Next = &ListNode{Val: 1}
			} else {
				pre.Next = &ListNode{Val: v}
				pre = pre.Next
			}
		} else {
			pre.Next = l1
			break
		}
		l1 = l1.Next
	}

	for l2 != nil {
		if pre.Next != nil {
			v := l2.Val + 1
			if v >= 10 {
				pre.Next = &ListNode{Val: v - 10}
				pre = pre.Next
				pre.Next = &ListNode{Val: 1}
			} else {
				pre.Next = &ListNode{Val: v}
				pre = pre.Next
			}
		} else {
			pre.Next = l2
			break
		}
		l2 = l2.Next
	}
	return temp.Next
}

func Test_addTwoNumbers(t *testing.T) {
	l1 := createNode(9)
	l2 := createNode(4)
	//l3 := createNode(3)
	l1.Next = l2
	//l2.Next = l3

	n1 := createNode(5)
	n2 := createNode(6)
	n3 := createNode(9)
	n1.Next = n2
	n2.Next = n3

	printListNode(addTwoNumbers(l1, n1))
}
