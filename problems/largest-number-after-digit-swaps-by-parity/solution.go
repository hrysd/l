func largestInteger(num int) int {
	even := []int{}
	odd := []int{}
	s := strconv.Itoa(num)
	for num > 0 {
		n := num % 10

		if n%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}

		num /= 10
	}

	sort.Ints(even)
	sort.Ints(odd)

	answer := 0

	for i := len(s) - 1; i >= 0; i-- {
		digit := s[i] - '0'

		var d int
		if digit%2 == 0 {
			d = even[0]
			even = even[1:]
		} else {
			d = odd[0]
			odd = odd[1:]
		}

		answer += d * int(math.Pow10(len(s)-i-1))
	}
	return answer
}
