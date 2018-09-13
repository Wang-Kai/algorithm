package sort

import (
	"testing"
)

func TestRandomizedSelect(t *testing.T) {
	arr := []int{12, 2, 2, 5, 4, 3, 45}

	val := RandomizedSelect(arr, 0, len(arr)-1, 0)
	t.Log(val)
}
