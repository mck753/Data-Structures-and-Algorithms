package medium

func findAnagrams(s string, p string) []int {
	l, r := 0, 0
	ret := make([]int, 0)
	have := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]] += 1
	}

	isOK := 0
	for r < len(s) {
		rChar := s[r]
		r++
		if _, ok := need[rChar]; ok {
			have[rChar] += 1
			if have[rChar] == need[rChar] {
				isOK++
			}
		}

		if r-l >= len(p) {
			if isOK == len(need) {
				ret = append(ret, l)
			}

			lChar := s[l]
			if _, ok := need[lChar]; ok {
				if have[lChar] == need[lChar] {
					isOK--
				}
				have[lChar] -= 1
			}
			l++
		}

	}

	return ret
}
