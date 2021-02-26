package main

import (
	"fmt"
	"math"
)

func readMatrix(length int) [][]int {
	matrix := make([][]int, length)
	for i := 0; i < length; i++ {
		matrix[i] = make([]int, length)
		for j := 0; j < length; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
	}
	return matrix
}

func main() {
	var length int
	fmt.Scanf("%d", &length)
	matrix := readMatrix(length)
	pdSum, sdSum := 0, 0

	for i := 0; i < length; i++ {
		pdSum += matrix[i][i]
		sdSum += matrix[length - 1 - i][i]
	}
	fmt.Println(math.Abs(float64(pdSum - sdSum)))
}
