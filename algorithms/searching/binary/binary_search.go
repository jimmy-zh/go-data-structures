package binary

import (
	"github.com/midnight-vivian/go-data-structures/utils"
)

func Search(arr []utils.Comparabler, item utils.Comparabler) int {
	if len(arr) == 0 {
		return 0
	}
	return searchRecursive(arr, item, 0, len(arr)-1)
}

func searchRecursive(arr []utils.Comparabler, item utils.Comparabler, lo, hi int) int {
	if lo > hi {
		return lo
	}
	m := (lo + hi) / 2
	if arr[m].Less(item) {
		return searchRecursive(arr, item, m+1, hi)
	} else if arr[m].More(item) {
		return searchRecursive(arr, item, lo, m-1)
	} else {
		return m
	}
}
