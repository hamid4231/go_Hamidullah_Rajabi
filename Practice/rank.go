package main

import "fmt"

func isPrime(n int) int {
	// Write your code here
	if n == 2 {
		return 1
	}
	if n%2 == 0 {
		return 2
	}
	for i := int(3); i <= n; i += 2 {
		if n%i == 0 {
			if i == n {
				return 1
			}
			return int(i)
		}
	}
	return 1
}

func main() {
	n := 4
	fmt.Println(isPrime(n))
}
