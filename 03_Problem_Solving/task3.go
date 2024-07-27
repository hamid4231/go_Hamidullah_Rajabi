package main

import (
	"fmt"
)

// Memoization map to store previously computed Fibonacci numbers
var memo = map[int]int{}

// Function to compute the n-th Fibonacci number using dynamic programming (top-down approach)
func fibonacci(n int) int {
	// Base cases
	if n <= 1 {
		return n
	}

	// Check if the result is already computed and stored in the memoization map
	if val, exists := memo[n]; exists {
		return val
	}

	// Recursively compute the Fibonacci number and store it in the memoization map
	memo[n] = fibonacci(n-1) + fibonacci(n-2)

	return memo[n]
}

func main() {
	// Example: Retrieve the 10th Fibonacci number
	n := 10
	fmt.Printf("The %d-th Fibonacci number is: %d\n", n, fibonacci(n))
}
