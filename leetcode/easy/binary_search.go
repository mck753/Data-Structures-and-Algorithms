package easy

func BinarySearch(arr []int, target int) int {
	return binarySearch(arr, target)
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
