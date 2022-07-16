func maxDistance(colors []int) int {
	length := len(colors)
	i, j := 0, length-1

	for colors[0] == colors[j] {
		j -= 1
	}

	for colors[length-1] == colors[i] {
		i += 1
	}

	if j > length-i-1 {
		return j
	} else {
		return length - i - 1
	}
}
