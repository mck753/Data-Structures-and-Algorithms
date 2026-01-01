package easy

func RemoveNthNode(head *ListNode, n int) *ListNode {
	return removeNthNode(head, n)
}

// 链表: 1->2->3->4->5，N=2 → 1->3->4->5
// N=1 → 2->3->4->5
func removeNthNode(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	curr := dummyNode
	for i := 1; i < n && curr != nil; i++ {
		curr = curr.Next
	}

	if curr == nil || curr.Next == nil {
		return head
	}

	curr.Next = curr.Next.Next

	return dummyNode.Next
}
