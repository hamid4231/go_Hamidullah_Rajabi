package main

import (
	"fmt"
	"slices"
)

func main() {
	// array works same as sql array which is a list of values
	// in go you have to specify the number of values inside the array

	// approach 1
	var number [5]int
	number[0] = 7
	number[1] = 10
	number[2] = 11
	number[3] = 8
	number[4] = 9
	// for i := 0; i <= len(number); i++ {
	// 	fmt.Println(number[i])
	// }
	for idx, val := range number {
		fmt.Println("the index: ", idx, " the value: ", val)
	}

	//slice is same as array but without number of values limit
	//approach 1
	var books []string
	books = append(books, "Data science", "Data Analysis", "Art")
	for idx, val := range books {
		fmt.Println("the index: ", idx, " the book: ", val)
	}
	//if we use under score (_) instead of idx, then index of the value will not be shown
	//reverse for loop
	for i := len(books) - 1; i >= 0; i-- {
		fmt.Println(books[i])
	}
	//sum of array
	var score []int = []int{12, 13, 14, 15}
	var sum int = 0
	for _, score := range score {
		sum += score
	}
	fmt.Println("the sum of array: ", sum)
	//slice package
	//checking if an item exist in an array
	var score []int = []int{12, 13, 14, 15}
	exist := slices.Contains(score, 5)
	fmt.Println(exist)
	//compact is used to remove duplicates
}
