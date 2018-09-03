package sort

import "testing"

func TestQuickSort(t *testing.T) {
	var arr = []int{204, 34, 20, -1, 489, 0, 9, 03, 4, 22343, 4, 23, 4, 0}

	QuickSort(arr)

	t.Logf("%+v", arr)
}
