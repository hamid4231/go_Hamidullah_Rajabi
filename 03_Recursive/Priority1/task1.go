package main

import "fmt"

// Function to calculate the sum of Fibonacci numbers up to a given limit
func sumFibonacci(limit int) int {
	// Handle the case where the limit is less than 1
	if limit < 1 {
		return 0
	}

	// Initial Fibonacci numbers
	a, b := 0, 1
	sum := 0

	// Loop to generate Fibonacci numbers and sum them until the limit is reached
	for a <= limit {
		sum += a
		a, b = b, a+b
	}

	return sum
}

func main() {
	// Test case 1
	limit := 5
	fmt.Printf("The sum of Fibonacci numbers up to %d is %d\n", limit, sumFibonacci(limit))
}
