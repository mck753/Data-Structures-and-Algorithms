package easy

// moveZeroes https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	cnt := 0
	for _, num := range nums {
		if num != 0 {
			nums[cnt] = num
			cnt++
		}
	}

	for i := cnt; i < len(nums); i++ {
		nums[i] = 0
	}
}
