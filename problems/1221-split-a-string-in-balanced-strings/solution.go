func balancedStringSplit(s string) int {
	splitted := strings.Split(s, "")
	count := 0
	result := 0

	for _, v := range splitted {
		if v == "R" {
			count++
		} else if v == "L" {
			count--
		}

		if count == 0 {
			result++
		}
	}

	return result
}
