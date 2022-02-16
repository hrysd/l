// append(nums, nums...)
func getConcatenation(nums []int) []int {
	result := make([]int, len(nums)*2)

	for i := 0; i < len(nums); i++ {
		result[i] = nums[i]
		result[len(nums)+i] = nums[i]
	}

	return result
}
