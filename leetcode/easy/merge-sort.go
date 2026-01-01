package easy

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	mid := l + (r-l)/2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	mergeArr(arr, l, mid, r)
}

func mergeArr(arr []int, l int, mid int, r int) {
	i, j := l, mid+1
	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			i++
		} else {
			tmp := arr[j]
			idx := j
			for idx > i {
				arr[idx] = arr[idx-1]
				idx--
			}

			arr[i] = tmp
			i++
			mid++
			j++
		}
	}
}
