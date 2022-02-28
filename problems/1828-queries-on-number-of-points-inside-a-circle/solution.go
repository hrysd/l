func countPoints(points [][]int, queries [][]int) []int {
	result := make([]int, len(queries))
	for i, q := range queries {
		x := q[0]
		y := q[1]
		r := q[2]
		rr := r * r
		count := 0
		for _, p := range points {
			x2 := p[0]
			y2 := p[1]

			xx := (x - x2)
			yy := (y - y2)

			if (xx*xx)+(yy*yy) <= rr {
				count++
			}
		}
		result[i] = count
	}

	return result
}
