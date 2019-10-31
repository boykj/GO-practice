package main

import "fmt"

func main() {

	sum(500, 12340, 45300)

}

func sum(input ...int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}

	fmt.Println("sum is ", sum)
	return sum
}

