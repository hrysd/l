func calPoints(ops []string) int {
	result := []int{}

	for _, o := range ops {
		if o == "+" {
			l := len(result)
			result = append(result, result[l-2]+result[l-1])
		} else if o == "D" {
			result = append(result, result[len(result)-1]*2)
		} else if o == "C" {
			result = result[:len(result)-1]
		} else {
			n, _ := strconv.Atoi(o)
			result = append(result, n)
		}
	}
	fmt.Println(result)
	return sum(result)
}

func sum(nums []int) int {
	count := 0

	for _, n := range nums {
		count += n
	}

	return count
}
