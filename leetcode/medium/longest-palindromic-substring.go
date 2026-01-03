package medium

func longestPalindrome(s string) string {
	n := len(s)
	start, maxLen := 0, 1
	expand := func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			// cbbd
			if r-l+1 > maxLen {
				start = l
				maxLen = r - l + 1
			}
			l--
			r++
		}
	}
	for i := 0; i < n; i++ {
		expand(i, i)
		expand(i, i+1)
	}

	return s[start : start+maxLen]
}
