package search

func BinarySearch(arr []int, target int) int {
	len := len(arr)
	low := 0
	high := len - 1
	result := -1

	boundary := 0
	for {

		if target == arr[high] {
			result = high
			break
		}

		if low == high {
			break
		}

		boundary = low + (high-low)/2
		if target > arr[boundary] {
			low = boundary + 1
		} else {
			high = boundary
		}
	}

	return result
}
