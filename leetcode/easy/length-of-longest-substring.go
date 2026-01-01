package easy

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	l, maxLen := 0, 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if idx, ok := m[s[i]]; ok && idx >= l {
			l = idx + 1
		}

		m[s[i]] = i
		maxLen = max(maxLen, i-l+1)
	}

	return maxLen
}
