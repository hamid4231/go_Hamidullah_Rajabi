package main

import (
	"fmt"
	"sync"
)

// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to generate prime numbers from 1 to 100 and send them to a channel
func generatePrimes(primeCh chan<- int) {
	defer close(primeCh)
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			primeCh <- i
		}
	}
}

// Function to calculate the square of numbers received from a channel and send them to another channel
func calculateSquares(primeCh <-chan int, squareCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for prime := range primeCh {
		squareCh <- prime * prime
	}
}

func main() {
	primeCh := make(chan int)
	squareCh := make(chan int)
	var wg sync.WaitGroup

	// Start goroutine to generate prime numbers
	go generatePrimes(primeCh)

	// Start goroutines to calculate squares
	for i := 0; i < 3; i++ { // Using 3 goroutines for calculating squares
		wg.Add(1)
		go calculateSquares(primeCh, squareCh, &wg)
	}

	// Close the squareCh channel once all squares are calculated
	go func() {
		wg.Wait()
		close(squareCh)
	}()

	// Print the squares of prime numbers
	for square := range squareCh {
		fmt.Println(square)
	}
}
