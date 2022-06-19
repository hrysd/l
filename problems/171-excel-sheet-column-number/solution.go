func titleToNumber(columnTitle string) int {
	number := 0

	for _, c := range columnTitle {
		number *= 26
		number += int(c-'A') + 1
	}

	return number
}
