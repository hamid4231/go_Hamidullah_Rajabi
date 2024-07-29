package main

import (
	"fmt"
	"time"
)

// Function to print each stage of the word reversal
func printReversedStage(word string) {
	reversed := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversed = string(word[i]) + reversed
		fmt.Println(reversed)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	word := "phone"
	printReversedStage(word)
}
