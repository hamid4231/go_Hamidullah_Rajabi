package main

import "fmt"

func main() {
	var bscore int
	var dscore int
	var difscore int
	var FinalScore int
	var budget int
	var duration int
	var difficulty int

	fmt.Print("Enter Budget Score: ")
	fmt.Scan(&budget)

	fmt.Print("Enter Duration Score: ")
	fmt.Scan(&duration)

	fmt.Print("Enter Difficulty Score: ")
	fmt.Scan(&difficulty)

	switch {
	case budget >= 0 && budget <= 50:
		bscore = 4
	case budget >= 51 && budget <= 80:
		bscore = 3
	case budget >= 81 && budget <= 100:
		bscore = 2
	case budget > 100:
		bscore = 1
	}

	switch {
	case duration >= 0 && duration <= 20:
		dscore = 10
	case duration >= 21 && duration <= 30:
		dscore = 5
	case duration >= 31 && duration <= 50:
		dscore = 2
	case duration > 50:
		dscore = 1
	}

	switch {
	case difficulty >= 0 && difficulty <= 3:
		difscore = 10
	case difficulty >= 4 && difficulty <= 6:
		difscore = 5
	case difficulty >= 7 && difficulty <= 10:
		difscore = 2
	case difficulty > 10:
		difscore = 1
	}

	FinalScore = bscore + dscore + difscore
	switch {
	case FinalScore <= 24 && FinalScore >= 17:
		fmt.Println("High")
	case FinalScore <= 16 && FinalScore >= 10:
		fmt.Println("Medium")
	case FinalScore <= 9 && FinalScore >= 3:
		fmt.Println("Low")
	case FinalScore <= 2:
		fmt.Println("Impossible")
	}
}
