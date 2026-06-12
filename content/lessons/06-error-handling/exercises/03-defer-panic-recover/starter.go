package main

import "fmt"

var _ = fmt.Println

// TODO: Write a function called "safeDivide" that takes two int parameters (a, b)
// and returns result int.
// 1. Use a deferred function with recover() to catch panics and print "Recovered: <value>".
// 2. If b is zero, call panic("division by zero").
// 3. Otherwise, return a / b.

func main() {
	// TODO: Call safeDivide(10, 0) and print the result as "10 / 0 = <result>"
	// TODO: Call safeDivide(10, 2) and print the result as "10 / 2 = <result>"
}
