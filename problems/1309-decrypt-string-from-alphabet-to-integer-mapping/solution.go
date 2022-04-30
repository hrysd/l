func freqAlphabets(s string) string {
	regex := regexp.MustCompile(`\d\d#|\d`)
	numbers := regex.FindAllString(s, -1)
	result := make([]rune, len(numbers))

	for i, number := range numbers {
		number = strings.ReplaceAll(number, "#", "")
		n, _ := strconv.Atoi(number)
		result[i] = rune(n + 'a' - 1)
	}

	return string(result)
}
