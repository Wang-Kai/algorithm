package sort

import "testing"

func Test(t *testing.T) {
	arr := []int{1, 0, 0, 2345, 345435, 567, 4, 4, 4, 4, 4, 90324, 4, 4, 4, 42, 45, 29}

	HeapSort(arr)

	t.Log(arr)
}
