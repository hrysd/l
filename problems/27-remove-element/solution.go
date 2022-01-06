func removeElement(nums []int, val int) int {
	length := len(nums)
	deleted := 0
	i := 0

	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			deleted++
		} else {
			i++
		}
	}

	return length - deleted
}
