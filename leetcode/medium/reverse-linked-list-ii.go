package medium

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left > right {
		return nil
	}

	dummyNode := &ListNode{Next: head}
	prev := dummyNode
	for i := 1; i < left; i++ {
		prev = prev.Next // 找出第 left-1 个
	}

	cur := prev.Next
	var newHead *ListNode
	for i := left; i <= right; i++ {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}

	// 循环结束，cur：right+1；newHead：新开头
	prev.Next.Next = cur
	prev.Next = newHead

	return dummyNode.Next
}
