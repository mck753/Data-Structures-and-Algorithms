package medium

func openLock(deadends []string, target string) int {
	// 记录需要跳过的死亡密码
	deads := make(map[string]struct{})
	for _, s := range deadends {
		deads[s] = struct{}{}
	}
	if _, found := deads["0000"]; found {
		return -1
	}

	// 记录已经穷举过的密码，防止走回头路
	visited := make(map[string]struct{})
	q := make([]string, 0)
	// 从起点开始启动广度优先搜索
	step := 0
	q = append(q, "0000")
	visited["0000"] = struct{}{}

	for len(q) > 0 {
		sz := len(q)
		// 将当前队列中的所有节点向周围扩散
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]

			// 判断是否到达终点
			if cur == target {
				return step
			}

			// 将一个节点的合法相邻节点加入队列
			for _, neighbor := range getNeighbors(cur) {
				if _, found := visited[neighbor]; !found {
					if _, dead := deads[neighbor]; !dead {
						q = append(q, neighbor)
						visited[neighbor] = struct{}{}
					}
				}
			}
		}
		// 在这里增加步数
		step++
	}
	// 如果穷举完都没找到目标密码，那就是找不到了
	return -1
}

// 将 s[j] 向上拨动一次
func plusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j]++
	}
	return string(ch)
}

// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j]--
	}
	return string(ch)
}

// 将 s 的每一位向上拨动一次或向下拨动一次，8 种相邻密码
func getNeighbors(s string) []string {
	neighbors := make([]string, 0)
	for i := 0; i < 4; i++ {
		neighbors = append(neighbors, plusOne(s, i))
		neighbors = append(neighbors, minusOne(s, i))
	}
	return neighbors
}
