package main

import (
	"fmt"

	sort "github.com/deankinane/algo-course/src/sort/quicksort"
)

func main() {
	data := []int{5, 8, 6, 0, 5, 2, 1, 12, 45, 79, 23, 14, 21, 25, 88, 3, 78, 4, 22, 10}

	sort.Quicksort(&data)
	fmt.Println(data)
}
