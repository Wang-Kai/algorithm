package sort

/*
	counting sort
	最坏运行时间：𝚯(n)
*/
func CountingSort(arr []int) []int {
	var tmpSlice = make([]int, 475969)

	for _, item := range arr {
		tmpSlice[item] = tmpSlice[item] + 1
	}

	for i := 1; i < len(tmpSlice); i++ {
		tmpSlice[i] = tmpSlice[i-1] + tmpSlice[i]
	}

	var res = make([]int, len(arr)+1)

	for j := len(arr) - 1; j >= 0; j-- {
		res[tmpSlice[arr[j]]] = arr[j]
		tmpSlice[arr[j]] = tmpSlice[arr[j]] - 1
	}

	return res[1:]
}
