package main

import (
	"base/leetcode/easy"
	"fmt"
)

func main() {
	r := findAnagrams("abab", "ab")
	fmt.Println(r)
}

func findAnagrams(s string, p string) []int {
	l, r := 0, 0
	ret := make([]int, 0)
	have := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]] += 1
	}

	isOK := 0
	// baa ，aa
	for r < len(s) {
		rChar := s[r]
		r++
		if _, ok := need[rChar]; ok {
			have[rChar] += 1
			if have[rChar] == need[rChar] {
				isOK++
			}
		}

		for r-l > len(p) {
			lChar := s[l]
			if _, ok := need[lChar]; ok {
				have[lChar] -= 1
				if have[lChar] < need[lChar] {
					isOK--
				}
			}
			l++
		}

		if isOK == len(need) {
			ret = append(ret, l)
		}

	}

	return ret
}

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := make(map[byte]int)
	required := 0 // 满足条件的字符种类数
	minLen := len(s) + 1
	minStart := -1
	l := 0

	for r := 0; r < len(s); r++ {
		c := s[r]
		if _, ok := need[c]; ok {
			have[c]++
			if have[c] == need[c] {
				required++
			}
		}

		// 当窗口覆盖所有字符时，收缩左指针
		for required == len(need) && l <= r {
			// 更新最小窗口
			if r-l+1 < minLen {
				minLen = r - l + 1
				minStart = l
			}

			// 移除左指针字符
			lc := s[l]
			if _, ok := need[lc]; ok {
				have[lc]--
				if have[lc] < need[lc] {
					required--
				}
			}
			l++
		}
	}

	if minStart == -1 {
		return ""
	}
	return s[minStart : minStart+minLen]
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
