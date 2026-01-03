package easy

func removeElement(nums []int, val int) int {
	slow := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			slow++
			nums[slow] = nums[i]
		}

	}

	return slow + 1
}
