func sortArrayByParityII(nums []int) []int {
	i, j := 0, 1
	for i < len(nums) && j < len(nums) {
		if nums[i]%2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j += 2
		} else {
			i += 2
		}
	}

	return nums
}
