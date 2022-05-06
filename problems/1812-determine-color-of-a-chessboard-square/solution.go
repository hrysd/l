func squareIsWhite(coordinates string) bool {
	column := coordinates[0] - 'a' + 1
	row := coordinates[1] - '0'

	return (column+row)%2 != 0
}
