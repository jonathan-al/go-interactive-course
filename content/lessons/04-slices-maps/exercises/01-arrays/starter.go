package main

import "fmt"

func main() {
	// TODO: Declare an array of 5 integers with values 10, 20, 30, 40, 50
	var numbers [5]int

	// TODO: Print the length of the array
	fmt.Println("Length:", len(numbers))

	// TODO: Print each element on a separate line
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
