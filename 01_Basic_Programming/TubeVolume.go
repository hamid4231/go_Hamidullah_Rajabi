package main

import "fmt"

func main() {
	var radius int
	var height int
	var volume float64
	pi := 3.14
	fmt.Print("Enter The Radius: ")
	fmt.Scan(&radius)
	fmt.Print("Enter The Height: ")
	fmt.Scan(&height)
	volume = pi * float64(radius*radius) * float64(height)
	fmt.Println("The volume of your tube is: ", volume)
}
