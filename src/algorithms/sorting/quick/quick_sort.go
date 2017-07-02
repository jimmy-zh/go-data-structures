package quick

const CUTOFF = 8

func Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	lo := 0
	hi := len(arr) - 1
	sort(arr, lo, hi)
}

func sort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	//median-of-3 for partitioning
	n := hi - lo + 1
	if n > CUTOFF {
		m := median3(arr, lo, hi, lo+n/2)
		arr[lo], arr[m] = arr[m], arr[lo]
	}
	p := partition(arr, lo, hi)
	sort(arr, lo, p-1)
	sort(arr, p+1, hi)
}

func median3(arr []int, i, j, k int) int {
	if arr[i] < arr[j] {
		if arr[j] < arr[k] {
			return j
		} else {
			if arr[i] < arr[k] {
				return k
			} else {
				return i
			}
		}
	} else {
		if arr[j] > arr[k] {
			return j
		} else {
			if arr[i] > arr[k] {
				return k
			} else {
				return i
			}
		}
	}
}

func partition(arr []int, lo, hi int) int {
	compare := arr[lo]
	i := lo
	j := hi + 1
	for {
		for {
			i++
			if i > hi || arr[i] >= compare {
				break
			}
		}
		for {
			j--
			if j <= lo || arr[j] <= compare {
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
