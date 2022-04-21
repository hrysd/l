func flipAndInvertImage(image [][]int) [][]int {
	length := len(image)

	for _, v := range image {
		for i := 0; i < length/2; i++ {
			v[i], v[length-i-1] = v[length-i-1], v[i]
		}

		for i, n := range v {
			v[i] = n ^ 1
		}
	}

	return image
}
