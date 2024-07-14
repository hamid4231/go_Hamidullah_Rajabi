package main

import (
	"fmt"
	"sort"
)

func sum(slice []int) float64 {
	total := 0
	for _, num := range slice {
		total += num
	}
	return float64(total)
}

func mean(slice []int) float64 {
	return sum(slice) / float64(len(slice))
}

func median(slice []int) float64 {
	sort.Ints(slice)
	n := len(slice)
	if n%2 == 1 {
		return float64(slice[n/2])
	}
	return float64(slice[n/2-1]+slice[n/2]) / 2.0
}

func mode(slice []int) float64 {
	frequency := make(map[int]int)
	maxFreq := 0
	modes := []int{}

	for _, num := range slice {
		frequency[num]++
		if frequency[num] > maxFreq {
			maxFreq = frequency[num]
		}
	}

	for num, freq := range frequency {
		if freq == maxFreq {
			modes = append(modes, num)
		}
	}

	sort.Ints(modes)
	return float64(modes[0])
}

func main() {
	data1 := []int{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Println("Dataset 1")
	fmt.Printf("sum: %.2f\n", sum(data1))
	fmt.Printf("mean: %.2f\n", mean(data1))
	fmt.Printf("median: %.2f\n", median(data1))
	fmt.Printf("mode: %.2f\n", mode(data1))

	data2 := []int{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Println("Dataset 2")
	fmt.Printf("sum: %.2f\n", sum(data2))
	fmt.Printf("mean: %.2f\n", mean(data2))
	fmt.Printf("median: %.2f\n", median(data2))
	fmt.Printf("mode: %.2f\n", mode(data2))
}
