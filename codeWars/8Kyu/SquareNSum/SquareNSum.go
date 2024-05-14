package main

import (
	"fmt"
)

// Complete the square sum function so that it squares each number passed into it and then sums the results together.
// For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2+ 2^2 = 9

func main() {
	fmt.Println(SquareSum([]int{1, 2, 2}))
}

func SquareSum(numbers []int) (resp int) {
	for _, n := range numbers {
		resp += n * n
	}
	return resp
}
