package sort

import "testing"

func TestCountingSort(t *testing.T) {
	var arr = []int{324, 345, 4, 76, 43, 5, 43, 80, 534, 5, 2, 5, 6, 47, 5, 60, 0, 0, 0, 0, 0, 7}
	result := CountingSort(arr)

	t.Logf("%+v", result)
}
