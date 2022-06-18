func destCity(paths [][]string) string {
	from := map[string]int{}

	for _, p := range paths {
		from[p[0]] = 1
	}

	for _, p := range paths {
		if _, ok := from[p[1]]; !ok {
			return p[1]
		}
	}

	return ""
}
