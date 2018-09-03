package sort

/*
	quick sort
	平均运行时间：𝚯(n * lgn)
*/
func QuickSort(arr []int) {
	from, to := 0, len(arr)-1
	if from < to {
		q := partition(arr, from, to)
		QuickSort(arr[from:q])
		QuickSort(arr[q+1 : to+1])
	}
}

func partition(arr []int, from, to int) int {
	x := arr[to]

	i := from - 1

	for j := from; j < to; j++ {
		if arr[j] < x {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[i+1], arr[to] = arr[to], arr[i+1]

	return i + 1
}
