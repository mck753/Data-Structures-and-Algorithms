package hard

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
