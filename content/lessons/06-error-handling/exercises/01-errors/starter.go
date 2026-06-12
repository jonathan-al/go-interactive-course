package main

import (
	"errors"
	"fmt"
)

var _ = errors.New
var _ = fmt.Println

// TODO: Write a function called "divide" that takes two int parameters (a, b)
// and returns (int, error). If b is zero, return an error with the message
// "cannot divide by zero". Otherwise, return a/b and nil.

func main() {
	// TODO: Call divide(10, 0), check for error, and print "Error: <message>"
	// TODO: Call divide(10, 3), and print the result as "10 / 3 = 3"
}
