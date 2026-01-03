package medium

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	have := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]] += 1
	}

	r, l, validCnt := 0, 0, 0
	for r < len(s2) {
		rChar := s2[r]
		r++
		if _, ok := need[rChar]; ok {
			have[rChar] += 1
			if need[rChar] == have[rChar] {
				validCnt += 1
			}
		}

		if r-l == len(s1) {
			if validCnt == len(need) {
				return true
			}

			lChar := s2[l]
			l++
			if _, ok := need[lChar]; ok {
				if need[lChar] == have[lChar] {
					validCnt -= 1
				}

				have[lChar] -= 1
			}
		}
	}

	return false
}
