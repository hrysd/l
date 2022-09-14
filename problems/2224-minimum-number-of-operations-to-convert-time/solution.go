func convertTime(current string, correct string) int {
	a := strings.Split(current, ":")

	ah, _ := strconv.Atoi(a[0])
	am, _ := strconv.Atoi(a[1])

	b := strings.Split(correct, ":")

	bh, _ := strconv.Atoi(b[0])
	bm, _ := strconv.Atoi(b[1])

	diff := (bh*60 + bm) - (ah*60 + am)

	answer := 0

	for _, operation := range []int{60, 15, 5, 1} {
		answer += diff / operation
		diff %= operation
	}
	return answer
}
