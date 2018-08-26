package main

import (
	"fmt"
)

func main() {
	var arr = []int{51, 30, 235, 75, 9, 24, 42, 486, 0, 190}

	// InsertionSort(arr)
	MergeSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

/*
	Merger sort
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

//-------------------------------------------------------

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
