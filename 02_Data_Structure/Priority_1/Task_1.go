package main

import "fmt"

func main() {
	res1 := merge([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(res1)
	res2 := merge([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(res2)
	res3 := merge([][]int{
		{1, 4, 5},
		{7},
	})

	fmt.Println(res3)
}

func merge(arr [][]int) []int {
	uniqueMap := make(map[int]bool)
	var result []int

	for _, subArray := range arr {
		for _, elem := range subArray {

			if _, exists := uniqueMap[elem]; !exists {

				uniqueMap[elem] = true
				result = append(result, elem)
			}
		}
	}

	return result
}
