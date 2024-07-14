package main

import (
	"errors"
	"fmt"
)

// func main() {
// 	var a, b int
// 	fmt.Print("Insert first number: ")
// 	fmt.Scan(&a)

// 	fmt.Print("Insert second number: ")
// 	fmt.Scan(&b)
// 	//first approach
// 	//sum(a,b)
// 	//second approach
// 	result := sum(a, b)
// 	fmt.Println("the sum result: ", result)
// }

// func sum(a int, b int) int {
// 	result := a + b
// 	//first returning result approach
// 	//fmt.Printf("%d+%d=%d/n", a,b,result)
// 	//second returning result approach
// 	return result
// }

// multiple returns
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}
func main() {
	res, err := div(12, 20)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("result: ", res)
	}
}
