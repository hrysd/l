func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return getNumericalValue(firstWord)+getNumericalValue(secondWord) == getNumericalValue(targetWord)
}

func getNumericalValue(word string) int {
	length := len(word)
	count := 0

	for i, char := range word {
		n := char - 'a'

		count += int(n) * int(math.Pow10(length-(i+1)))
	}

	return count
}
