func areNumbersAscending(s string) bool {
	splitted := strings.Split(s, " ")
	current := 0

	for _, v := range splitted {
		if n, err := strconv.Atoi(v); err == nil {
			if current < n {
				current = n
			} else {
				return false
			}
		}
	}

	return true
}
