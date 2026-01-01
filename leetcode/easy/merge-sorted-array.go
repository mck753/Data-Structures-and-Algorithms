package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	j, k := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if j < 0 {
			nums1[i] = nums2[k]
			k--
		} else if k < 0 {
			nums1[i] = nums1[j]
			j--
		} else if nums1[j] > nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
	}
}
