package search

func BubbleSort(data []int) []int {
	max := len(data)

	for max > 1 {
		for i := 0; i < max-1; i++ {
			if data[i] > data[i+1] {
				swap := data[i]
				data[i] = data[i+1]
				data[i+1] = swap
			}
		}
		max--
	}

	return data
}
