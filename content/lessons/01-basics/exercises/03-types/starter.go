package main

import "fmt"

func main() {
	a := 42
	b := 3.14
	c := "Hello"
	d := true

	// TODO: Print each variable's type and value using fmt.Printf
	// Use %T for type and %v for value
	// Expected format: "name: type=TYPE, value=VALUE"
	//
	// Example:
	//   fmt.Printf("a: type=%T, value=%v\n", a, a)
	//   produces: a: type=int, value=42

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
