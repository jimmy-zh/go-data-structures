package bubble

func Sort(arr []int) {
	if len(arr) < 2 {
		return
	}
	swap := true
	for i := 0; i < len(arr)-1 && swap; i++ {
		swap = false
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				swap = true
			}
		}
	}
}
