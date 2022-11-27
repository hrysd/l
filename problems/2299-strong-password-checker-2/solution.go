func strongPasswordCheckerII(password string) bool {
	lower := false
	upper := false
	digit := false
	special := false

	if len(password) < 8 {
		return false
	}

	for _, char := range password {
		if 'a' <= char && char <= 'z' {
			lower = true
		} else if 'A' <= char && char <= 'Z' {
			upper = true
		} else if '0' <= char && char <= '9' {
			digit = true
		} else {
			special = true
		}
	}

	if !(lower && upper && digit && special) {
		return false
	}

	for i := 1; i < len(password); i++ {
		char := password[i]

		if password[i-1] == char {
			return false
		}
	}

	return true
}


