func findOcurrences(text string, first string, second string) []string {
	answer := []string{}
	splitted := strings.Split(text, " ")

	for i := 0; i < len(splitted)-2; i++ {
		if splitted[i] == first && splitted[i+1] == second {
			answer = append(answer, splitted[i+2])
		}
	}

	return answer
}
