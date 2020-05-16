package 递归

import (
	"fmt"
	"testing"
)

/*
https://leetcode-cn.com/explore/orignial/card/recursion-i/260/conclusion/1229/


{ list1[0]+merge(list1[1:],list2)    list1[0]<list2[0]
{ list2[0]+merge(list1,list2[1:])     otherwise
​
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func Test_mergeTwoLists(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3

	n4 := &ListNode{Val: 1}
	n5 := &ListNode{Val: 3}
	n6 := &ListNode{Val: 4}
	n4.Next = n5
	n5.Next = n6

	r := mergeTwoLists(n1, n4)
	for r != nil {
		fmt.Print(r.Val, " ")
		r = r.Next
	}

}
