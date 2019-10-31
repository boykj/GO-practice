package main

import "fmt"

func main() {

	sum(500, 12340, 45300)

}

// Basic sum function - will accept variadic arguments and sum all integeres included within function call
func sum(input ...int) int {
	sum := 0
	for _, i := range input {
		fmt.Println("sum is >>> ", sum, "value is >>>", i)
		sum += i
	}

	fmt.Println("sum is ", sum)
	return sum
}

