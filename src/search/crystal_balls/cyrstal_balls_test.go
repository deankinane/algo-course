package crystalballs

import "testing"

func TestCrystalBalls(t *testing.T) {
	data := []bool{}

	for i := 0; i < 100; i++ {
		if i < 69 {
			data = append(data, false)
		} else {
			data = append(data, true)
		}
	}

	result := CrystalBalls(data)

	if result != 69 {
		t.Error("Expected 69, got ", result)
	}
}
