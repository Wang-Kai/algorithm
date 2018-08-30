package sort

/*
	Merger sort
	最坏运行时间：𝚯(n * lgn)
*/
func MergeSort(arr []int, from, to int) {
	if from < to {
		halfTag := (from + to) / 2
		MergeSort(arr, from, halfTag)
		MergeSort(arr, halfTag+1, to)
		merge(arr, from, halfTag, to)
	}
}

/*
	Merge sort
*/

// merge 是归并排序的一个子程序，负责排序数组中已经排序好的两段数组
// [p, q]是已经排序好的一段数字，(q, r] 是一段已经排序好的数字，index 从 0 开始
func merge(arr []int, p, q, r int) {
	leftArr := make([]int, q-p+1)
	rightArr := make([]int, r-q)

	copy(leftArr, arr[p:q+1])
	copy(rightArr, arr[q+1:])

	var i, j int

	for k := p; k <= r; k++ {
		if leftArr[i] > rightArr[j] {
			arr[k] = rightArr[j]
			j++

			if j >= len(rightArr) {
				// 将 leftArr 中剩余的项全部填到 arr 尾部，结束循环
				copy(arr[k+1:], leftArr[i:])
				break
			}
		} else {
			arr[k] = leftArr[i]
			i++

			if i >= len(leftArr) {
				copy(arr[k+1:], rightArr[j:])
				break
			}
		}
	}
}
