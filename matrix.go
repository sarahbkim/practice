package main

func spiralOrder(matrix [][]int) []int {
	var rowStart, colStart int
	var rowEnd, colEnd = len(matrix) - 1, len(matrix[rowStart]) - 1

	ans := spiralRowCol(matrix, rowStart, rowEnd, colStart, colEnd)

	return ans
}

func spiralRowCol(matrix [][]int, rowStart, rowEnd, colStart, colEnd int) []int {
	if rowStart > rowEnd || colStart > colEnd {
		return nil
	}
	var ans []int
	for i := colStart; i <= colEnd; i++ {
		ans = append(ans, matrix[rowStart][i])
	}
	for i := rowStart + 1; i <= rowEnd; i++ {
		ans = append(ans, matrix[i][colEnd])
	}
	if rowEnd != rowStart {
		for i := colEnd - 1; i >= colStart; i-- {
			ans = append(ans, matrix[rowEnd][i])
		}
	}
	if colEnd != colStart {
		for i := rowEnd - 1; i > rowStart; i-- {
			ans = append(ans, matrix[i][colStart])
		}
	}

	ans = append(ans, spiralRowCol(matrix, rowStart+1, rowEnd-1, colStart+1, colEnd-1)...)
	return ans
}

func rotate(matrix [][]int) {
	rotateCols(matrix)
	rotateDiag(matrix)
}

func rotateCols(matrix [][]int) {
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		for col := 0; col < len(matrix[i]); col++ {
			matrix[i][col], matrix[j][col] = matrix[j][col], matrix[i][col]
		}
	}
}

func rotateDiag(matrix [][]int) {
	var j int
	for row := 0; row < len(matrix); row++ {
		for col := j; col < len(matrix); col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
		j += 1
	}
}
