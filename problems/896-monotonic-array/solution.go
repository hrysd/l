func isMonotonic(nums []int) bool {
	ascending := true
	descending := true

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			descending = false
		}

		if nums[i-1] < nums[i] {
			ascending = false
		}
	}

	return ascending || descending
}
