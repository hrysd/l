func totalMoney(n int) int {
	savings := 0
	d := n % 7
	w := n / 7
	savings += w * 28

	if w > 0 {
		savings += ((w*(w+1))/2)*7 - ((7 - d) * w)
	}

	for i := 1; i <= d; i++ {
		savings += i
	}

	return savings
}
