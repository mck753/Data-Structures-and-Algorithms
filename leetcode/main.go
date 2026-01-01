package main

import (
	"base/leetcode/easy"
	"fmt"
)

func main() {
	// 长度为 5 的数组
	arr := []int{1, 2, 3, 4, 5}
	i := 0
	// 模拟环形数组，这个循环永远不会结束
	for i < len(arr) {
		fmt.Println(arr[i])
		i = (i + 1) % len(arr)
	}
}

func buildList() *easy.ListNode {
	return &easy.ListNode{1, &easy.ListNode{2, &easy.ListNode{3, &easy.ListNode{4, &easy.ListNode{5, nil}}}}}
}

func printList(head *easy.ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func mergeArr(arr []int, l, mid, r int) {
	i, j := l, mid+1
	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			i++
		} else {
			tmp := arr[j]
			idx := j
			for idx > i {
				arr[idx] = arr[idx-1]
				idx--
			}

			arr[i] = tmp

			i++
			mid++
			j++
		}
	}
}
