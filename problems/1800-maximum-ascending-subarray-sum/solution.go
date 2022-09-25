func maxAscendingSum(nums []int) int {
	count := nums[0]
	tmp := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			tmp += nums[i]
		} else {
			tmp = nums[i]
		}

		if count < tmp {
			count = tmp
		}
	}

	return count
}
