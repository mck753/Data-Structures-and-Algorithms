package hard

import "base/leetcode/medium"

func reverseKGroup(head *medium.ListNode, k int) *medium.ListNode {
	if head == nil || k == 1 {
		return head
	}

	// 哑节点，方便处理
	dummy := &medium.ListNode{Next: head}
	pre := dummy

	for pre.Next != nil {
		// 检查剩余数量是否充足
		cur := pre.Next
		cnt := 0
		// 1>2>3,2
		for cur != nil && cnt < k {
			cur = cur.Next
			cnt++
		}
		if cnt < k {
			break
		}

		newHead, newTail := reverse(pre.Next, k)
		pre.Next = newHead
		pre = newTail
	}

	return dummy.Next
}

func reverse(head *medium.ListNode, k int) (*medium.ListNode, *medium.ListNode) {
	if head == nil {
		return nil, nil
	}

	cur := head
	//tail := head
	var prev *medium.ListNode
	for i := 0; i < k; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	head.Next = cur

	return prev, head
}
