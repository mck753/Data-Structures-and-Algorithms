package medium

func partition(head *ListNode, x int) *ListNode {
	l1 := &ListNode{}
	l2 := &ListNode{}
	curl1 := l1
	curl2 := l2
	cur := head
	for cur != nil {
		if cur.Val < x {
			curl1.Next = cur
			curl1 = curl1.Next
		} else {
			curl2.Next = cur
			curl2 = curl2.Next
		}

		tmp := cur.Next
		cur.Next = nil
		cur = tmp
	}

	if l1.Next != nil {
		curl1.Next = l2.Next
		return l1.Next
	}

	return l2.Next
}
