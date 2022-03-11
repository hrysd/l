func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	count := 0
	rules := map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}

	index := rules[ruleKey]

	for _, item := range items {
		if item[index] == ruleValue {
			count++
		}
	}

	return count
}
