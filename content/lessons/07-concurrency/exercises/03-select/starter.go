package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "first"
		ch2 <- "second"
	}()

	// TODO: Use a for loop with 2 iterations.
	// Inside the loop, use a select statement with two cases:
	//   case msg := <-ch1: print msg
	//   case msg := <-ch2: print msg

	_ = ch1
	_ = ch2
	_ = fmt.Println
}
