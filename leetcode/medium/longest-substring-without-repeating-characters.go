package medium

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	if len(runes) <= 1 {
		return len(runes)
	}

	l, r := 0, 0
	ret := 0
	have := make(map[rune]int)
	for r < len(runes) {
		rChar := runes[r]
		have[rChar] += 1
		r++

		for have[rChar] > 1 && l < r {
			lChar := runes[l]
			l++
			have[lChar] -= 1
		}
		if r-l > ret {
			ret = r - l
		}
	}

	return ret
}
