package sort

import (
	"fmt"
	"testing"
)

func TestQuicksort(t *testing.T) {
	data := []int{5, 8, 6, 0, 5, 2, 1, 12, 45, 79, 23, 14, 21, 25, 88, 3, 78, 4, 22, 10}
	solution := "[0 1 2 3 4 5 5 6 8 10 12 14 21 22 23 25 45 78 79 88]"

	Quicksort(&data)

	if fmt.Sprint(data) != solution {
		t.Errorf("Exptected %v but got %v", solution, data)
	}

}
