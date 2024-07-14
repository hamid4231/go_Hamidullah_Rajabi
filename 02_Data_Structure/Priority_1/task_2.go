package main

import (
	"fmt"
)

func primeSum(slice []int) int {
	sum := 0

	for _, num := range slice {
		if num <= 1 {
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
			sum += num
		}
	}

	return sum
}

func main() {
	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30}))
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))
	fmt.Println(primeSum([]int{15, 12, 9, 10}))
}
