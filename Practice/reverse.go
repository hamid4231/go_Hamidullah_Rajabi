package main

import (
	"fmt"
	"slices"
)

func main() {
	//get input
	var word string
	fmt.Print("Insert word: ")
	fmt.Scan(&word)

	//reverse using slices package
	var letters []byte = []byte(word)
	slices.Reverse(letters)
	var reversed string = string(letters)
	fmt.Println(reversed)
}
