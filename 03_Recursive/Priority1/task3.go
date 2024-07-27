package main

import "fmt"

// Function to calculate the n-th triangular number
func pattern(n int) int {
	return (n * (n + 1)) / 2
}

func main() {
	// Test Case 1
	input1 := 100
	output1 := pattern(input1)
	fmt.Printf("Input: %d\nOutput: %d\n", input1, output1)
}
