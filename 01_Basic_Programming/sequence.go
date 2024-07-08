package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		switch {
		case number%4 == 0 && number%7 == 0:
			fmt.Println("Buzz")
		case number%4 == 0:
			fmt.Println(number * number)
		case number%7 == 0:
			fmt.Println(number * number * number)
		default:
			fmt.Println(number)
		}
	}
}
