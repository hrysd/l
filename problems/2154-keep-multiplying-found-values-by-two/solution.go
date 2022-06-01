func findFinalValue(nums []int, original int) int {
	num := original
	found := true

	for found {
		found = false
		for _, n := range nums {
			if n == num {
				num = num * 2
				found = true
				break
			}
		}
	}

	return num
}
