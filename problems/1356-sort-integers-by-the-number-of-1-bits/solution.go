func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a := strings.Count(strconv.FormatInt(int64(arr[i]), 2), "1")
		b := strings.Count(strconv.FormatInt(int64(arr[j]), 2), "1")

		if a == b {
			return arr[i] < arr[j]
		} else {
			return a < b
		}

	})

	return arr
}
