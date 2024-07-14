package main

import (
	"fmt"
	"math"
)

// isPrime function checks if a given number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrtN; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(100007))
	fmt.Println(isPrime(13))
	fmt.Println(isPrime(17))
	fmt.Println(isPrime(20))
	fmt.Println(isPrime(35))
}
