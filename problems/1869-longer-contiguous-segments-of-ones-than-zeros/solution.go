func checkZeroOnes(s string) bool {
	one := strings.Split(s, "0")
	zero := strings.Split(s, "1")

	sortByLength(&one)
	sortByLength(&zero)

	return len(one[0]) > len(zero[0])
}

func sortByLength(list *[]string) {
	sort.Slice(*list, func(i, j int) bool {
		return len((*list)[i]) > len((*list)[j])
	})
}
