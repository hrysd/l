func shuffle(nums []int, n int) []int {
	result := make([]int, n*2)
	index := 0

	for i, j := 0, n; i < n; i, j = i+1, j+1 {
		result[index] = nums[i]
		index += 1
		result[index] = nums[j]
		index += 1
	}

	return result
}
