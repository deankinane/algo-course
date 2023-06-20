package main

import (
	"fmt"

	search "github.com/deankinane/algo-course/src/search/bubble_sort"
)

func main() {

	data := []int{2, 1, 8, 5, 7, 3, 4, 9, 6}

	data = search.BubbleSort(data)
	fmt.Println(data)
}
