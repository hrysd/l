func kLengthApart(nums []int, k int) bool {
	appearedIndex := -k - 1

	for i, v := range nums {
		if v != 1 {
			continue
		}

		if i-appearedIndex <= k {
			return false
		}

		appearedIndex = i
	}

	return true
}
