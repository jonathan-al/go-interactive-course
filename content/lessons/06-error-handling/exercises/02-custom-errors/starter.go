package main

import (
	"errors"
	"fmt"
	"math"
)

var _ = errors.As
var _ = fmt.Println
var _ = math.Sqrt

// TODO: Define a struct called "NegativeError" with a field "Value int".

// TODO: Implement the Error() string method on *NegativeError so it returns
// "negative value: <Value>" (e.g., "negative value: -4").

// TODO: Write a function called "sqrt" that takes a float64 and returns (float64, error).
// If the input is negative, return &NegativeError{Value: int(input)} and an error.
// Otherwise, return math.Sqrt(input) and nil.

func main() {
	// TODO: Call sqrt(-4). Use errors.As to check for *NegativeError and print:
	// "NegativeError: negative value: -4"

	// TODO: Call sqrt(16) and print the result as "sqrt(16) = 4.000000"
}
