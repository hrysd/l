func arraySign(nums []int) int {
	sign := 1

	for _, n := range nums {
		if n > 0 {
			n = 1
		} else if n < 0 {
			n = -1
		} else {
			n = 0
		}

		sign *= n

		if sign == 0 {
			break
		}
	}

	return sign
}
