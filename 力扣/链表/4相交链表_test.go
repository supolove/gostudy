package 链表

import (
	"fmt"
	"testing"
)

/*
	https://leetcode-cn.com/explore/learn/card/linked-list/194/two-pointer-technique/746/
*/

// 496 ms O(n^2)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	for curA != nil {
		curB = headB
		//curAA := curA
		for curB != nil {
			//if curAA == curB {
			//	curAA = curAA.Next
			//	if curAA == nil || curB.Next == nil{
			//		return curA
			//	}
			//}
			if curA == curB {
				return curA
			}
			curB = curB.Next
		}
		curA = curA.Next
	}
	return nil
}

// 用map实现
func getIntersectionNode2(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	m := map[*ListNode]bool{}
	curA := headA
	for curA != nil {
		m[curA] = true
		curA = curA.Next
	}

	curB := headB
	for curB != nil {
		_, ok := m[curB]
		if ok {
			return curB
		}
		curB = curB.Next
	}
	return nil
}

// 两个连表互相连接相同距离同时到达相同点
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p0 := headA
	p1 := headB
	first := true
	for p0 != p1 {
		if p0.Next == nil {
			if first {
				p0 = headB
				first = false
			} else {
				break
			}
		} else {
			p0 = p0.Next
		}
		if p1.Next == nil {
			p1 = headA
		} else {
			p1 = p1.Next
		}
	}
	if p0 != p1 {
		return nil
	}
	return p0
}

// 32 毫秒
func getIntersectionNode4(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB

	var pc *ListNode = nil

	if pa == nil || pb == nil {
		return pc
	}

	n := 0
	for n < 3 {
		if pa == nil {
			n++
			pa = headB
			continue
		}

		if pb == nil {
			n++
			pb = headA
			continue
		}

		if pa != pb {
			pa = pa.Next
			pb = pb.Next
		} else {
			pc = pa
			break
		}
	}
	return pc
}

func Test_getIntersectionNode(t *testing.T) {
	node6 := &ListNode{Val: 8}
	node7 := &ListNode{Val: 4}
	node8 := &ListNode{Val: 5}
	node6.Next = node7
	node7.Next = node8

	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 1}
	node1.Next = node2
	node2.Next = node6

	node3 := &ListNode{Val: 5}
	node4 := &ListNode{Val: 0}
	node5 := &ListNode{Val: 1}
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	fmt.Println(getIntersectionNode3(node1, node3))

}
