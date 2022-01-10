func plusOne(digits []int) []int {
	result := []int{}
	length := len(digits) - 1
	advanced := false

	for j := length; j >= 0; j-- {
		d := digits[j]

		if length == j {
			if d+1 == 10 {
				advanced = true
				result = append([]int{0}, result...)

			} else {
				result = append([]int{d + 1}, result...)
			}
			continue
		}

		if !advanced {
			result = append([]int{d}, result...)
			continue
		}

		if d+1 == 10 {
			advanced = true
			result = append([]int{0}, result...)
		} else {
			advanced = false
			result = append([]int{d + 1}, result...)
		}
	}

	if advanced {
		result = append([]int{1}, result...)
	}

	return result
}
