func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maximum int

	for _, v := range candies {
		if v > maximum {
			maximum = v
		}
	}

	result := make([]bool, len(candies))

	for i, v := range candies {
		if v+extraCandies >= maximum {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}
