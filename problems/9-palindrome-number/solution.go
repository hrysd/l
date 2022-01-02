func isPalindrome(x int) bool {
	str := fmt.Sprint(x)
	strRune := []rune(str)

	start := 0
	end := len(strRune) - 1

	for start < end {
		strRune[start], strRune[end] = strRune[end], strRune[start]
		start++
		end--
	}

	return string(strRune) == str
}
