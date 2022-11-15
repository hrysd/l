func countHillValley(nums []int) int {
	count := 0
	left := nums[0]

	for i := 0; i < len(nums)-1; i++ {
		valley := left < nums[i] && nums[i] > nums[i+1]
		hill := left > nums[i] && nums[i] < nums[i+1]

		if valley || hill {
			count += 1
			left = nums[i]
		}
	}

	return count
}
