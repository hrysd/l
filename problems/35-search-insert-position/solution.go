func searchInsert(nums []int, target int) int {
	left := -1
	right := len(nums)

	for right-left > 1 {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid
		} else {
			left = mid
		}

	}

	return right
}
