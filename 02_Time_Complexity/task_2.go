package main

import (
	"fmt"
)

// power function calculates x raised to the power n using exponentiation by squaring
func power(x int, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	result := 1
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}
	return result
}

func main() {
	fmt.Println(power(2, 3))
	fmt.Println(power(5, 3))
	fmt.Println(power(10, 2))
	fmt.Println(power(2, 5))
	fmt.Println(power(7, 3))
}
