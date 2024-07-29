package main

import (
	"fmt"
	"sync"
)

// Function to check if a number is composite
func isComposite(n int) bool {
	if n <= 1 {
		return false
	}
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count > 2
}

// Function to generate composite numbers from 1 to 100 and send them to a channel
func generateComposites(compositeCh chan<- int) {
	defer close(compositeCh)
	for i := 1; i <= 100; i++ {
		if isComposite(i) {
			compositeCh <- i
		}
	}
}

// Function to calculate the square of numbers received from a channel and send them to another channel
func calculateSquares(compositeCh <-chan int, squareCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for composite := range compositeCh {
		squareCh <- composite * composite
	}
}

// Function to check if a number is even or odd and print "Pong" or "Ping" accordingly
func checkEvenOdd(squareCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for square := range squareCh {
		if square%2 == 0 {
			fmt.Println("Pong")
		} else {
			fmt.Println("Ping")
		}
	}
}

func main() {
	compositeCh := make(chan int)
	squareCh := make(chan int)
	var wg sync.WaitGroup

	// Start goroutine to generate composite numbers
	go generateComposites(compositeCh)

	// Start goroutines to calculate squares
	for i := 0; i < 3; i++ { // Using 3 goroutines for calculating squares
		wg.Add(1)
		go calculateSquares(compositeCh, squareCh, &wg)
	}

	// Close the squareCh channel once all squares are calculated
	go func() {
		wg.Wait()
		close(squareCh)
	}()

	// Start a goroutine to check if the squares are even or odd
	wg.Add(1)
	go checkEvenOdd(squareCh, &wg)

	// Wait for the final goroutine to complete
	wg.Wait()
}
