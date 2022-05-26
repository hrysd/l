func replaceElements(arr []int) []int {
	length := len(arr)

	for i, _ := range arr {
		if i == length-1 {
			arr[i] = -1
		} else {
			arr[i] = max(arr[i+1:])
		}
	}

	return arr
}

func max(arr []int) int {
	max := 0

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}
