package easy

func QuickSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	quickSort(arr, 0, n-1)
}

func quickSort(arr []int, l int, r int) {
	if l < r {
		p := partition(arr, l, r)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, r)
	}
}

func partition(arr []int, l int, r int) int {
	p := arr[r]
	i := l - 1
	for j := l; j < r; j++ {
		if arr[j] < p {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
