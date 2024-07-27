package main

import (
	"fmt"
)

// Function to generate Pascal's Triangle
func generatePascalTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

// Function to print Pascal's Triangle
func printPascalTriangle(triangle [][]int) {
	for _, row := range triangle {
		fmt.Println(row)
	}
}

func main() {
	numRows := 5 // Change this value to generate a different number of rows
	triangle := generatePascalTriangle(numRows)
	printPascalTriangle(triangle)
}
