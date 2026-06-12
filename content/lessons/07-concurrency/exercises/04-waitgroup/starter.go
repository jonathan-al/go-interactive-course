package main

import (
	"fmt"
	"sync"
)

type Result struct {
	ID    int
	Value int
}

func main() {
	// TODO: Declare a sync.WaitGroup

	// TODO: Create a buffered channel of type Result with capacity 5

	// TODO: Launch 5 goroutines (for numbers 1 to 5).
	// Each goroutine should:
	//   - defer wg.Done()
	//   - compute n*n*n (cube of n)
	//   - send Result{ID: n, Value: n*n*n} to the channel
	// Use: go func(n int) { ... }(i)

	// TODO: Launch a goroutine that waits for the WaitGroup and closes the channel:
	// go func() { wg.Wait(); close(ch) }()

	// TODO: Receive all results using for range and store them in a slice
	// at index result.ID - 1.

	// TODO: Print each result in order using:
	// fmt.Printf("ID: %d, Value: %d\n", result.ID, result.Value)

	_ = fmt.Println
	_ = sync.WaitGroup{}
}
