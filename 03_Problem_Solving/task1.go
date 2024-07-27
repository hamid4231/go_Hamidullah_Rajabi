package main

import (
	"fmt"
	"strconv"
)

func intToBinary(n int) []string {
	result := make([]string, n+1)
	for i := 0; i <= n; i++ {
		binary := strconv.FormatInt(int64(i), 2)
		result[i] = binary
	}
	return result
}

func main() {
	n := 3 // You can change this value to test other inputs
	binaries := intToBinary(n)
	fmt.Println(binaries) // Output: [0, 1, 10, 11]
}
