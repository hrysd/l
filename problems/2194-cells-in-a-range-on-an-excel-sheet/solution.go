func cellsInRange(s string) []string {
	result := []string{}
	cells := strings.Split(s, ":")

	start := cells[0]
	end := cells[1]

	scol := start[0]
	srow := start[1]

	ecol := end[0]
	erow := end[1]

	for c := scol; c <= ecol; c++ {
		for r := srow; r <= erow; r++ {
			result = append(result, fmt.Sprintf("%s%s", string(c), string(r)))
		}
	}

	return result
}
