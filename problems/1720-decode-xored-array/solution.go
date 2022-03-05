func decode(encoded []int, first int) []int {
	result := make([]int, len(encoded)+1)
	result[0] = first

	for i, v := range encoded {
		result[i+1] = result[i] ^ v
	}

	return result
}
