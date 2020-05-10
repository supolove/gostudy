package 链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func createSimpleListNode() *ListNode {
	node := createNode(1)
	node2 := createNode(2)
	node3 := createNode(3)
	node4 := createNode(2)
	node5 := createNode(1)

	node.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	return node
}

func printListNode(list *ListNode) {
	c := list

	for i := 0; c != nil; i++ {
		print(c.Val, " ")
		c = c.Next
		if i > 100 {
			println()
			println("超出100个元素")
			break
		}
	}
	println()
}

func createNode(val int) *ListNode {
	return &ListNode{Val: val}
}
