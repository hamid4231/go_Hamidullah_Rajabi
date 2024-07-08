package main

import (
	"fmt"
)

// Function to check if two words are anagrams
func areAnagrams(word1, word2 string) bool {
	// If the lengths of the words are different, they cannot be anagrams
	if len(word1) != len(word2) {
		return false
	}

	// Create a map to count the frequency of each letter in the first word
	letterCount := make(map[rune]int)

	// Count the frequency of each letter in the first word
	for _, letter := range word1 {
		letterCount[letter]++
	}

	// Decrease the count for each letter in the second word
	for _, letter := range word2 {
		letterCount[letter]--
		// If the count goes below zero, the words are not anagrams
		if letterCount[letter] < 0 {
			return false
		}
	}

	// If all counts are zero, the words are anagrams
	for _, count := range letterCount {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	var word1, word2 string

	fmt.Println("Enter the first word: ")
	fmt.Scanln(&word1)

	fmt.Println("Enter the second word: ")
	fmt.Scanln(&word2)

	if areAnagrams(word1, word2) {
		fmt.Println("Anagram")
	} else {
		fmt.Println("Not an anagram")
	}
}
