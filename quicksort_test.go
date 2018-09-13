package sort

import "testing"

func TestQuickSort(t *testing.T) {
	var arr = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

	QuickSort(arr)

	t.Logf("%+v", arr)
}
