package 链表

import "testing"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	curl1 := l1
	curl2 := l2
	curl2Pre := l2
OuterLoop:
	for curl1 != nil {
		temp := curl1
		for curl2 != nil {
			if curl1.Val <= curl2.Val {
				temp = curl1.Next
				// 是否头元素
				if curl2.Val == l2.Val {
					curl1.Next = l2
					l2 = curl1
					curl2Pre = l2
				} else {
					// 在中间插入元素
					curl1.Next = curl2
					curl2Pre.Next = curl1
					curl2Pre = curl1
				}
				break
			} else {

				if curl2.Next == nil {
					// 结尾元素拼接到2链表
					curl2.Next = curl1
					break OuterLoop
				}
				// curl2 +1
				curl2Pre = curl2
				curl2 = curl2.Next

			}
		}
		curl1 = temp
	}
	return l2
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}

	if l1 == nil {
		pre.Next = l2
	}
	if l2 == nil {
		pre.Next = l1
	}

	return dummy.Next

}

func Test_mergeTwoLists(t *testing.T) {
	/*
		l1 := createNode(1)
		l2 := createNode(3)
		l3 := createNode(5)
		l4 := createNode(11)
		l1.Next = l2
		l2.Next = l3
		l3.Next = l4

		n1 := createNode(1)
		n2 := createNode(6)
		n3 := createNode(8)
		n1.Next = n2
		n2.Next = n3

		printListNode(mergeTwoLists(l1, n1))

	*/

	l1 := createNode(-30)
	l2 := createNode(-30)
	l3 := createNode(5)
	l4 := createNode(11)
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

	n1 := createNode(-30)
	n2 := createNode(-30)
	n3 := createNode(6)
	n4 := createNode(8)
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	printListNode(mergeTwoLists2(l1, n1))

}
