package shell

func Sort(arr []int) {
	if len(arr) < 2 {
		return
	}
	n := len(arr)
	h := 1
	//find proper increment sequence
	for h < n/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < len(arr); i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h = h / 3
	}
}
