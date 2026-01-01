package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList https://leetcode.cn/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseListV2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
}

func majorityElement(nums []int) int {
	ret, cnt := 0, 0
	for _, num := range nums {
		if cnt == 0 {
			ret = num
			cnt = 1
		} else if ret == num {
			cnt++
		} else if ret != num {
			cnt--
		}
	}

	return ret
}

func maxProfit(prices []int) int {
	minPrice, ret := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > ret {
			ret = prices[i] - minPrice
		}
	}

	return ret
}

func romanToInt(s string) int {
	mapping := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	runes := []rune(s)
	if len(runes) == 1 {
		return mapping[runes[0]]
	}

	ret := 0
	for i := 0; i < len(runes)-1; i++ {
		cur, next := mapping[runes[i]], mapping[runes[i+1]]
		if cur < next { // 特殊case
			if cur == 1 {
				if next == 5 {
					ret += 4
				}
				if next == 10 {
					ret += 9
				}
			} else if cur == 10 {
				if next == 50 {
					ret += 40
				}
				if next == 100 {
					ret += 90
				}
			} else if cur == 100 {
				if next == 500 {
					ret += 400
				}
				if next == 1000 {
					ret += 900
				}
			}
			i++
		} else {
			ret += cur
		}
	}

	if mapping[runes[len(runes)-2]] > mapping[runes[len(runes)-1]] {
		ret += mapping[runes[len(runes)-1]]
	}

	return ret
}

func lengthOfLastWord(s string) int {
	cur, ret := 0, 0
	for _, elem := range s {
		if elem == ' ' && cur == 0 {
			continue
		}
		if elem == ' ' {
			ret = cur
			cur = 0
		} else {
			cur++
		}
	}

	if cur == 0 {
		return ret
	}

	return cur
}

func longestCommonPrefix(strs []string) string {
	ret := []rune(strs[0])
	for i := 1; i < len(strs); i++ {
		cur := []rune(strs[i])
		minLength := min(len(ret), len(cur))
		for j := 0; j < minLength; j++ {
			if ret[j] != cur[j] {
				ret = cur[:j]
				break
			}
		}

		if len(ret) == 0 {
			return ""
		}

		if len(ret) > minLength {
			ret = ret[:minLength+1]
		}
	}

	return string(ret)
}
