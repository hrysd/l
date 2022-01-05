func removeDuplicates(nums []int) int {
	k := len(nums)
	i := 0

	if k == 0 {
		return k
	}

	for {
		if i == len(nums)-1 {
			break
		}

		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			k--
		} else {
			i++
		}
	}

	return k
}
