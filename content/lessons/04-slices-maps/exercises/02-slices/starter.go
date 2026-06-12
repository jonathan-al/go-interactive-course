package main

import "fmt"

func main() {
	// TODO: Create a slice with initial values 1, 2, 3
	numbers := []int{0}

	// TODO: Append 4 and 5 to the slice
	numbers = append(numbers, 0)

	// TODO: Print the slice
	fmt.Println("Slice:", numbers)

	// TODO: Create and print a sub-slice from index 1 to 3 (exclusive)
	sub := numbers[0:1]
	fmt.Println("Sub-slice:", sub)
}
