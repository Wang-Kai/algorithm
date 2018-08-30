package sort

/*
	Heap sort
	最坏运行时间：𝚯(n * lgn)
*/
func HeapSort(arr []int) {
	buildHeapify(arr)

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		arr = arr[:i]
		maxHeapify(arr, 0)
	}
}

func buildHeapify(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		maxHeapify(arr, i)
	}
}

func maxHeapify(arr []int, i int) {
	l, r := Left(i), Right(i)

	var largest = i
	if l < len(arr) && arr[l] > arr[i] {
		largest = l
	}

	if r < len(arr) && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest)
	}
}

// heap base operation
func Parent(i int) int {
	return i / 2
}
func Right(i int) int {
	return 2*i + 2
}
func Left(i int) int {
	return 2*i + 1
}
