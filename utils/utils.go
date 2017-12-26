package utils

import "math/rand"

func randArray(n int) (arr []int) {
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return
}

func SortTest(n int, sort func([]int)) bool {
	arr := randArray(n)
	sort(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
