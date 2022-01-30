import "math"

func isHappy(n int) bool {
	numbers := map[int]bool{}
	number := n

	for true {
		splitted := split(number)
		number = sum(splitted)

		_, exists := numbers[number]

		if number == 1 || exists {
			break
		} else {
			numbers[number] = true
		}
	}

	return number == 1
}

func split(n int) []int {
	var slice []int

	for n != 0 {
		slice = append(slice, n%10)
		n /= 10
	}

	return slice
}

func sum(n []int) int {
	var s int

	for _, v := range n {
		s += int(math.Pow(float64(v), 2))
	}

	return s
}
