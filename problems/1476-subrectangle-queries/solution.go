type SubrectangleQueries struct {
	Rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{Rectangle: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for row := row1; row <= row2; row += 1 {
		for col := col1; col <= col2; col += 1 {
			this.Rectangle[row][col] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.Rectangle[row][col]
}

