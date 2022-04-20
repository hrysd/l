func uniqueMorseRepresentations(words []string) int {
	morseCodes := map[string]int{}
	maps := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	for _, w := range words {
		morseCode := make([]string, len(w))

		for i, char := range w {
			morseCode[i] = maps[char-'a']
		}

		code := strings.Join(morseCode, "")

		_, ok := morseCodes[code]

		if !ok {
			morseCodes[code] = 1
		}
	}

	return len(morseCodes)
}
