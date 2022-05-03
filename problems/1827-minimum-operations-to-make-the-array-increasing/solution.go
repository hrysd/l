func minOperations(nums []int) int {
	count := 0

	for i, n := range nums {
		if i == 0 {
			continue
		}

		if !(n > nums[i-1]) {
			val := nums[i-1] + 1
			count += val - n
			nums[i] = val
		}
	}

	return count
}
