package 链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func createSimpleListNode() *ListNode {
	node := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node.Next = node2
	node2.Next = &ListNode{Val: 0, Next: &ListNode{Val: -4, Next: node2}}
	return node
}
