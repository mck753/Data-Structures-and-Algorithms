package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: &ListNode{},
		size: 0,
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}

	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size || index < 0 {
		return
	}

	l.size++
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	node := &ListNode{Val: val, Next: cur.Next}
	cur.Next = node
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index >= l.size || index < 0 {
		return
	}

	l.size--
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
}
