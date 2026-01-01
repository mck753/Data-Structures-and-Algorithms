package hard

func reversePairs(record []int) int {
	return mergeSort(record, 0, len(record)-1)
}

func mergeSort(arr []int, l int, r int) int {
	if l >= r {
		return 0
	}

	ret := 0
	mid := l + (r-l)/2
	ret += mergeSort(arr, l, mid)
	ret += mergeSort(arr, mid+1, r)
	ret += merge(arr, l, mid, r)

	return ret
}

func merge(arr []int, l int, mid int, r int) int {
	tmp := make([]int, r-l+1)
	copy(tmp, arr[l:])

	i, j := l, mid+1
	ret := 0
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = tmp[j-l]
			j++
		} else if j > r {
			arr[k] = tmp[i-l]
			i++
		} else if tmp[i-l] <= tmp[j-l] {
			arr[k] = tmp[i-l]
			i++
		} else {
			ret += mid - i + 1
			arr[k] = tmp[j-l]
			j++
		}
	}

	return ret
}
