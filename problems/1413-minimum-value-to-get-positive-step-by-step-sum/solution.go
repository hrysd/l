func minStartValue(nums []int) int {
	min := 0

	current := 0
	for _, n := range nums {
		current += n

		if min > current {
			min = current
		}
	}

	return min*-1 + 1
}
