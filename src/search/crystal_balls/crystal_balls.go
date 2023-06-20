package crystalballs

import "math"

func CrystalBalls(arr []bool) int {

	n := len(arr)
	sqrt := int(math.Floor(math.Sqrt(float64(n))))

	for i := 0; i < n; i += sqrt {
		if arr[i+sqrt] {
			for j := i; j < i+sqrt; j++ {
				if arr[j] {
					return j
				}
			}
		}
	}

	return -1
}
