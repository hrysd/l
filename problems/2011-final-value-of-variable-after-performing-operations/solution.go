func finalValueAfterOperations(operations []string) int {
	var result int

	for _, o := range operations {
		if strings.Contains(o, "+") {
			result++
		} else {
			result--
		}
	}

	return result
}
