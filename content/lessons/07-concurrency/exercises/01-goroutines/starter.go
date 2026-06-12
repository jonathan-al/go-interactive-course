package main

import "fmt"

type Result struct {
	ID    int
	Value int
}

func main() {
	// TODO: Create a buffered channel of type Result with capacity 5

	// TODO: Launch 5 goroutines (for numbers 1 to 5).
	// Each goroutine should compute n*n and send Result{ID: n, Value: n*n} to the channel.
	// Use: go func(n int) { ... }(i)

	// TODO: Receive all 5 results from the channel and store them in a slice
	// at index result.ID - 1.

	// TODO: Print each result in order using:
	// fmt.Printf("ID: %d, Value: %d\n", result.ID, result.Value)

	_ = fmt.Println
}
