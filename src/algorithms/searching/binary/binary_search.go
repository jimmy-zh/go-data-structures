package binary

type Comparabler interface {
	Len() int
	Less(i int, t interface{}) bool
	More(i int, t interface{}) bool
}

func Search(arr Comparabler, item interface{}) int {
	if arr.Len() == 0 {
		return 0
	}

	return searchRecursive(arr, item, 0, arr.Len()-1)
}

func searchRecursive(arr Comparabler, item interface{}, lo, hi int) int {
	if lo > hi {
		return lo
	}
	m := (lo + hi) / 2
	if arr.Less(m, item) {
		return searchRecursive(arr, item, m+1, hi)
	} else if arr.More(m, item) {
		return searchRecursive(arr, item, lo, m-1)
	} else {
		return m
	}
}
