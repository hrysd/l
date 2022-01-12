func longestCommonPrefix(strs []string) string {
	prefix := ""
	str := strs[0]

	if len(str) == 0 {
		return prefix
	}

	length := len(strs)

Exit:
	for charAt := 0; charAt < len(str); charAt++ {
		char := string(str[charAt])

		for i := 1; i < length; i++ {
			if len(strs[i])-1 < charAt {
				break Exit
			}

			otherChar := string(strs[i][charAt])

			if char != otherChar {
				break Exit
			}
		}

		prefix += char
	}

	return prefix
}
