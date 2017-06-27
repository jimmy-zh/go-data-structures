package shell

func Sort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	h := 1
	//find proper increment sequence
	for h < len(arr)/3 {
		h = h*3 + 1
	}

	for h >= 1 {

		for i := h; i < len(arr); i++ {
			for j := i; j >= h; j -= h {
				if arr[j] < arr[j-h] {
					arr[j], arr[j-h] = arr[j-h], arr[j]
				}
			}
		}

		h = h / 3
	}

}
