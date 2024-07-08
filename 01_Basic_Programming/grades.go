package main

import "fmt"

func main() {
	var score int
	fmt.Print("Enter your score: ")
	fmt.Scan(&score)

	if score >= 85 && score <= 100 {
		fmt.Println(score, " A")
	} else if score >= 70 && score <= 84 {
		fmt.Println(score, " B")
	} else if score >= 55 && score <= 69 {
		fmt.Println(score, " C")
	} else if score >= 40 && score <= 54 {
		fmt.Println(score, " D")
	} else if score >= 0 && score <= 39 {
		fmt.Println(score, "E")
	} else {
		fmt.Println(score, " Invalid Score!")
	}
}
