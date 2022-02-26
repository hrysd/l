func numJewelsInStones(jewels string, stones string) int {
	chars := strings.Split(jewels, "")
	count := 0
	for _, v := range chars {
		count += strings.Count(stones, v)
	}

	return count
}
