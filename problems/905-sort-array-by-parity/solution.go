func sortArrayByParity(nums []int) []int {
	j := 0
	for i, n := range nums {
		if n%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	return nums
}
