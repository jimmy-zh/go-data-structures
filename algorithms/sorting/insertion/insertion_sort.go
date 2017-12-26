package insertion

func Sort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
