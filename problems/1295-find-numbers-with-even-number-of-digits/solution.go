func findNumbers(nums []int) int {
	count := 0

	for _, n := range nums {
		b := strconv.Itoa(n)

		if len(b)%2 == 0 {
			count++
		}
	}

	return count
}
