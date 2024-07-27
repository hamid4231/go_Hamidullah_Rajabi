package main

import (
	"fmt"
	"math"
)

// Function to find the integers x, y, z
func findIntegers(A, B, C int) (int, int, int, bool) {
	// Limit the search to a reasonable range
	limit := int(math.Sqrt(float64(C))) + 1

	// Iterate over possible values for x, y, z
	for x := -limit; x <= limit; x++ {
		for y := -limit; y <= limit; y++ {
			for z := -limit; z <= limit; z++ {
				if x != y && y != z && x != z { // Ensure x, y, z are different
					// Check if the values satisfy the equations
					if x+y+z == A && x*y*z == B && x*x+y*y+z*z == C {
						return x, y, z, true
					}
				}
			}
		}
	}

	// If no solution is found, return false
	return 0, 0, 0, false
}

func main() {
	// Example input values
	A := 6
	B := 6
	C := 14

	// Find the integers x, y, z
	x, y, z, found := findIntegers(A, B, C)
	if found {
		fmt.Printf("The integers are: x = %d, y = %d, z = %d\n", x, y, z)
	}
