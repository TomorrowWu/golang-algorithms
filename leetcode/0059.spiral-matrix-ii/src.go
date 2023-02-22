package _059_spiral_matrix_ii

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			matrix[top][column] = num
			num++
		}
		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = num
			num++
		}
		for column := right - 1; column >= left; column-- {
			matrix[bottom][column] = num
			num++
		}
		for row := bottom - 1; row > top; row-- {
			matrix[row][left] = num
			num++
		}
		left++
		right--
		top++
		bottom--
	}
	return matrix
}
