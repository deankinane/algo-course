package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []int{2, 1, 8, 5, 7, 3, 4, 9, 6}

	data = BubbleSort(data)
	for i, v := range data {
		if i+1 != v {
			t.Errorf("Expected %v but got %v", i+1, v)
		}
	}
}
