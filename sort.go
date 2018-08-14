package main

import (
	"fmt"
)

func main() {
	var arr = []int{12, 3, 56, 7, 9, 24, 2, 2, 390, 42, 34}

	InsertionSort(arr)
	fmt.Println(arr)
}

// insert sort
func InsertionSort(params []int) {
	if len(params) == 1 {
		return
	}

	for i := 1; i < len(params); i++ {
		var key = params[i]

		j := i - 1

		for j >= 0 && params[j] > key {
			params[i] = params[j]
			i = j
			j--
		}

		params[j+1] = key
	}
}
