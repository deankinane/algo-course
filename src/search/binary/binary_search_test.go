package search

import "testing"

func TestBinarySearch(t *testing.T) {
	data := []int{}

	for i := 1; i <= 1000; i++ {
		if i == 300 {
			data = append(data, i-1)
			continue
		}
		data = append(data, i)
	}

	result := BinarySearch(data, 910)
	if result != 909 {
		t.Error("Expected 909, got ", result)
	}

	result = BinarySearch(data, 123)
	if result != 122 {
		t.Error("Expected 122, got ", result)
	}

	result = BinarySearch(data, 1111)
	if result != -1 {
		t.Error("Expected -1, got ", result)
	}

	result = BinarySearch(data, 300)
	if result != -1 {
		t.Error("Expected -1, got ", result)
	}
}
