package main

import (
	"fmt"
)

// Function to group prime numbers from a slice
func groupPrimes(slice []int) ([][]int, []int) {
	primes := []int{}
	nonPrimes := []int{}

	for _, num := range slice {
		if num <= 1 {
			nonPrimes = append(nonPrimes, num)
			continue
		}
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, num)
		} else {
			nonPrimes = append(nonPrimes, num)
		}
	}
	return [][]int{primes}, nonPrimes
}

func main() {

	fmt.Println(groupPrimes([]int{2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(groupPrimes([]int{15, 24, 3, 9, 5}))
	fmt.Println(groupPrimes([]int{4, 8, 9, 12}))
}
