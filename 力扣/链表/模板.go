package 链表

// 判断链表是否有环模板
func moban() bool {
	head := &ListNode{}
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
