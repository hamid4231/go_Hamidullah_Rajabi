package main

import "fmt"

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
}

func getInfo(students []Student) {
	bestMathStudent := students[0]
	bestScienceStudent := students[0]
	bestEnglishStudent := students[0]
	bestAverageStudent := students[0]

	for _, student := range students {
		if student.MathScore > bestMathStudent.MathScore {
			bestMathStudent = student
		}
		if student.ScienceScore > bestScienceStudent.ScienceScore {
			bestScienceStudent = student
		}
		if student.EnglishScore > bestEnglishStudent.EnglishScore {
			bestEnglishStudent = student
		}
		averageScore := (student.MathScore + student.ScienceScore + student.EnglishScore) / 3
		bestAverageScore := (bestAverageStudent.MathScore + bestAverageStudent.ScienceScore + bestAverageStudent.EnglishScore) / 3
		if averageScore > bestAverageScore {
			bestAverageStudent = student
		}
	}

	fmt.Printf("best student in math: %s with %d\n", bestMathStudent.Name, bestMathStudent.MathScore)
	fmt.Printf("best student in science: %s with %d\n", bestScienceStudent.Name, bestScienceStudent.ScienceScore)
	fmt.Printf("best student in english: %s with %d\n", bestEnglishStudent.Name, bestEnglishStudent.EnglishScore)
	averageScore := (bestAverageStudent.MathScore + bestAverageStudent.ScienceScore + bestAverageStudent.EnglishScore) / 3
	fmt.Printf("best student in average: %s with %d\n", bestAverageStudent.Name, averageScore)
}
