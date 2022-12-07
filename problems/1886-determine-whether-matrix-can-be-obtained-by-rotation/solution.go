func findRotation(mat [][]int, target [][]int) bool {
	count := 0

	for count <= 3 {
		if reflect.DeepEqual(mat, target) {
			return true
		}
		mat = rotate(mat)
		count += 1
	}

	return false
}

func rotate(matrix [][]int) [][]int {
	l := len(matrix[0])

	for i := 0; i < l/2; i++ {
		matrix[i], matrix[l-i-1] = matrix[l-i-1], matrix[i]
	}

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return matrix
}
