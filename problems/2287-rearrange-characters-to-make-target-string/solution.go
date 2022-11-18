func rearrangeCharacters(s string, target string) int {
	sm := map[rune]int{}
	tm := map[rune]int{}

	for _, c := range target {
		tm[c] += 1
	}

	for _, c := range s {
		sm[c] += 1
	}

	answer := math.MaxInt

	for k, v := range tm {
		count := sm[k] / v

		if count < answer {
			answer = count
		}
	}

	return answer
}
