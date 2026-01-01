package easy

func FindKthFromEnd(head *ListNode, k int) *ListNode {
	return findKthFromEnd(head, k)
}

func findKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}

	fast := head
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
