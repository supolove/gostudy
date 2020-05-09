package 链表

import (
	"testing"
)

/*
https://leetcode-cn.com/explore/learn/card/linked-list/195/classic-problems/753/
*/

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var evenHead *ListNode
	if cur.Next != nil {
		evenHead = cur.Next
	}

	curEven := evenHead
	for cur.Next != nil {
		if cur != head {
			curEven.Next = cur.Next
			curEven = curEven.Next
		}

		cur.Next = cur.Next.Next
		if cur.Next == nil {
			if cur.Next != nil {
				cur = cur.Next
			}
			break
		}
		curEven.Next = nil
		cur = cur.Next
	}
	cur.Next = evenHead
	return head
}

func oddEvenList2(head *ListNode) *ListNode {
	var oddHead, oddNext *ListNode
	var evenHead, evenNext *ListNode

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	oddHead, oddNext = head, head
	evenHead, evenNext = head.Next, head.Next
	head = head.Next.Next
	oddNext.Next = nil
	evenNext.Next = nil
	for head != nil { //3
		//fmt.Printf("odd: %v\n", head.Val)

		oddNext.Next = head
		oddNext = oddNext.Next
		head = head.Next   //4
		oddNext.Next = nil //1->3

		if head != nil {
			//fmt.Printf("even: %v\n", head.Val)

			evenNext.Next = head //2->4
			evenNext = evenNext.Next
			head = head.Next
			evenNext.Next = nil

		}
	}
	//println(evenHead)
	oddNext.Next = evenHead
	return oddHead
}

func Test_oddEvenList(t *testing.T) {
	node := createSimpleListNode()
	printListNode(oddEvenList2(node))
}
