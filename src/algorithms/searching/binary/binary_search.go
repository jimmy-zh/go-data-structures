package binary

type Comparabler interface {
	Less(item interface{}) bool
	More(item interface{}) bool
}

func Search(arr []Comparabler, item Comparabler) int {
	if len(arr) == 0 {
		return 0
	}
	return searchRecursive(arr, item, 0, len(arr)-1)
}

func searchRecursive(arr []Comparabler, item Comparabler, lo, hi int) int {
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
