func findMiddleIndex(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	before := 0
	after := sum

	for i, n := range nums {
		if before == after-n {
			return i
		}

		before += n
		after -= n
	}

	return -1
}
