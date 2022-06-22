func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sSplitted := strings.Split(s, "")
	tSplitted := strings.Split(t, "")

	sort.Strings(sSplitted)
	sort.Strings(tSplitted)

	for i, c := range sSplitted {
		if c != tSplitted[i] {
			return false
		}
	}

	return true
}
