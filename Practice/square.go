package main

import "fmt"

func main() {
	var number int
	fmt.Print("How many sequence of number do you want to have the squared? ")
	fmt.Scan(&number)

	for i:= 1, i>= number, i++{
		fmt.Println(i*i)
	}
}