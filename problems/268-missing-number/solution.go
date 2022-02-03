func missingNumber(nums []int) int {
	sum := len(nums)

	for i, v := range nums {
		sum = sum ^ v
		sum = sum ^ i
	}

	return sum
}
