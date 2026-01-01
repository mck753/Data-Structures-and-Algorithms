package easy

// KthToLast https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/
func kthToLast(head *ListNode, k int) int {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		k--

		if k < 0 {
			slow = slow.Next
		}
	}

	return slow.Val
}
