func similarPairs(words []string) int {
	answer := 0

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if bitmask(words[i]) == bitmask(words[j]) {
				answer += 1
			}
		}
	}

	return answer
}

func bitmask(w string) int {
	mask := 0
	for _, c := range w {
		mask |= 1 << (c - 'a')
	}
	return mask
}
