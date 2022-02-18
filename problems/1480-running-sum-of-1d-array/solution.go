func runningSum(nums []int) []int {
	count := 0
	result := make([]int, len(nums))

	for i, v := range nums {
		count += v
		result[i] = count
	}

	return result
}
