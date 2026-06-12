package main

import "fmt"

func main() {
	var name string = "Alice"
	var age int = 30
	var height float64 = 1.68
	var student bool = true

	// TODO: Print each variable in the format "Label: value"
	// Expected output:
	//   Name: Alice
	//   Age: 30
	//   Height: 1.68
	//   Student: true
	//
	// Hint: fmt.Println("Name:", name) prints "Name: Alice"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(student)
}
