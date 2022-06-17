func buildArray(target []int, n int) []string {
	result := []string{}
	i := 0
	count := 1

	for i < n || count < n {
		if i == len(target) {
			break
		}

		if count == target[i] {
			result = append(result, "Push")
			i++
			count++
		} else {
			result = append(result, "Push")
			result = append(result, "Pop")
			count++
		}

	}
	return result
}
