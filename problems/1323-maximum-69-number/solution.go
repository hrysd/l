func maximum69Number(num int) int {
	appeared := false
	str := strconv.Itoa(num)

	result := make([]rune, len(str))

	for i, v := range str {
		if v == '6' && !appeared {
			result[i] = '9'
			appeared = true
		} else {
			result[i] = v
		}
	}

	n, _ := strconv.Atoi(string(result))

	return n
}
