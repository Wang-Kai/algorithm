package sort

/*
	insert sort
	最坏运行时间：𝚯(n * n)
*/
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
