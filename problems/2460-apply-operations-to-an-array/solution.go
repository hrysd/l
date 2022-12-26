func applyOperations(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			nums[i-1] = nums[i-1] * 2
			nums[i] = 0
		}
	}

	answer := []int{}

	for _, n := range nums {
		if n > 0 {
			answer = append(answer, n)
		}
	}

	return append(answer, make([]int, len(nums)-len(answer))...)
}
