package easy

import "math"

func maxAdjacentDistance(nums []int) int {
	var ret int
	for i := 0; i < len(nums)-1; i++ {
		ret = max(ret, int(math.Abs(float64(nums[i]-nums[i+1]))))
	}

	ret = max(ret, int(math.Abs(float64(nums[0]-nums[len(nums)-1]))))

	return ret
}
