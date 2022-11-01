func findSpecialInteger(arr []int) int {
	n := len(arr) / 4

	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[i+n] {
			return arr[i]
		}
	}

	return -1
}
