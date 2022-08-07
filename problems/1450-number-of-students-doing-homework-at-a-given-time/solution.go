func busyStudent(startTime []int, endTime []int, queryTime int) int {
	result := 0

	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			result += 1
		}
	}

	return result
}
