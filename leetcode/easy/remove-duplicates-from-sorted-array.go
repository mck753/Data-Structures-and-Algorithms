package easy

// removeDuplicates https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	cnt := 0
	for i, num := range nums {
		if i == 0 || num != nums[i-1] {
			nums[cnt] = num
			cnt++
		}
	}

	return cnt
}
