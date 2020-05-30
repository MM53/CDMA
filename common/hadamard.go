package common

func HadamardMatrix(order int) [][]int8 {
	if order == 1 {
		return [][]int8{{1}}
	} else {
		previousMatrix := HadamardMatrix(order - 1)
		previousSize := len(previousMatrix)
		var matrix [][]int8
		for y := 0; y < previousSize*2; y++ {
			// build row of matrix
			row := make([]int8, previousSize*2)
			for x := 0; x < previousSize*2; x++ {
				// copy value for field from previous matrix
				row[x] = previousMatrix[y%previousSize][x%previousSize]
				// multiply lower, right quarter with -1
				if x/previousSize == 1 && y/previousSize == 1 {
					row[x] *= -1
				}
			}

			matrix = append(matrix, row)
		}
		return matrix
	}
}
