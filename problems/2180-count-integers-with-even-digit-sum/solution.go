func countEven(num int) int {
	sum := 0
	tmp := num

	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	if sum%2 == 0 {
		return tmp / 2
	} else {
		return (tmp - 1) / 2
	}
}
