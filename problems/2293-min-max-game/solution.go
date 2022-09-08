func minMaxGame(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	answer := []int{}
	for i := 0; i < len(nums)/2; i++ {
		var n int
		a, b := nums[i*2], nums[i*2+1]

		if i%2 == 0 {
			if a < b {
				n = a
			} else {
				n = b
			}
		} else {
			if a > b {
				n = a
			} else {
				n = b
			}
		}
		answer = append(answer, n)
	}
	return minMaxGame(answer)
}
