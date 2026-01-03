package easy

// x>1>2
// x>2
// x
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Val: -1, Next: head}
	cur := dummyNode
	for i := 0; i < n+1; i++ {
		if cur == nil {
			return nil
		}

		cur = cur.Next
	}

	slow := dummyNode
	for cur != nil {
		cur = cur.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummyNode.Next
}
