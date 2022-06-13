func maximumUnits(boxTypes [][]int, truckSize int) int {
	result := 0
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })

	for _, box := range boxTypes {
		if truckSize == 0 {
			break
		}

		if truckSize > box[0] {
			result += box[0] * box[1]
			truckSize -= box[0]
		} else {
			result += truckSize * box[1]
			truckSize = 0
		}

	}

	return result
}
