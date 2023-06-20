package search

import "testing"

func TestLinearSearch(t *testing.T) {
	test := []int{1, 2, 3, 4, 5}
	result := LinearSearch(test, 3)
	if result != true {
		t.Error("Expected true, got ", result)
	}

	result = LinearSearch(test, 6)
	if result != false {
		t.Error("Expected false, got ", result)
	}
}
