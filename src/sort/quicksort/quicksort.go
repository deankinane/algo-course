package sort

func Quicksort(arr *[]int) {
	data := *arr
	qs(arr, 0, len(data)-1)
}

func qs(arr *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)
	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr *[]int, lo int, hi int) int {
	data := *arr
	p := data[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		if data[i] <= p {
			idx++
			tmp := data[i]
			data[i] = data[idx]
			data[idx] = tmp
		}
	}

	idx++
	tmp := data[idx]
	data[idx] = data[hi]
	data[hi] = tmp

	return idx
}
