package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter Your Number: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if number%i == 0 { // Check if i is a factor of number
			if i%2 == 0 {
				fmt.Print("I")
			} else {
				fmt.Print("O")
			}
		}
	}
	fmt.Println() // Print a newline at the end
}
