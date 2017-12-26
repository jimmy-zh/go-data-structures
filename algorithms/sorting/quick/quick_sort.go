package quick

import (
	"github.com/midnight-vivian/go-data-structures/src/algorithms/sorting/insertion"
)

const CUTOFF = 8

func Sort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	n := hi - lo + 1
	if n <= CUTOFF {
		insertion.Sort(arr[lo : hi+1])
		return
	}
	m := Median3(arr, lo, lo+n/2, hi)
	arr[lo], arr[m] = arr[m], arr[lo]

	i := partition(arr, lo, hi)
	sort(arr, lo, i-1)
	sort(arr, i+1, hi)
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[lo]
	i := lo
	j := hi + 1
	for {
		for {
			i++
			if arr[i] >= pivot || i == hi {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot || j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[lo], arr[j] = arr[j], arr[lo]
	return j
}

func Median3(arr []int, i, j, k int) int {
	if arr[i] < arr[j] {
		if arr[j] < arr[k] {
			return j
		}
		if arr[i] < arr[k] {
			return k
		}
		return i
	}
	if arr[i] < arr[k] {
		return i
	}
	if arr[j] < arr[k] {
		return k
	}
	return j
}
