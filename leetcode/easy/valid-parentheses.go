package easy

// isValid https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	left := make([]rune, 0)
	for _, elem := range s {
		if elem == '(' || elem == '{' || elem == '[' {
			left = append(left, elem)
			continue
		}

		if len(left) == 0 {
			return false
		}

		right := left[len(left)-1]
		if elem == ')' && right != '(' {
			return false
		} else if elem == '}' && right != '{' {
			return false
		} else if elem == ']' && right != '[' {
			return false
		}

		left = left[:len(left)-1]
	}

	return len(left) == 0
}
