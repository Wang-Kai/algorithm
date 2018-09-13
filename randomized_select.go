package sort

import (
	"math/rand"
	"time"
)

/*
	Randomized Select
	期望运行时间： 𝚯(n)
*/
func RandomizedSelect(arr []int, from, to, target int) int {
	rand.Seed(time.Now().Unix())

	if from == to {
		return arr[from]
	}

	q := randomizedPartition(arr, from, to)
	k := q - from

	if target == k {
		return arr[q]
	} else if target < k {
		return RandomizedSelect(arr, from, q-1, target)
	} else {
		return RandomizedSelect(arr, q+1, to, target-q-1)
	}
}

func randomizedPartition(arr []int, from, to int) int {
	randomIndex := rand.Intn(to-from+1) + from
	arr[to], arr[randomIndex] = arr[randomIndex], arr[to]

	return partition(arr, from, to)
}
