package main

import "fmt"

const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

const (
	StatusOK       = iota + 200 // 200
	StatusCreated               // 201
	StatusAccepted              // 202
)

func main() {
	// TODO: Print Sunday, Monday, and Tuesday with their iota values
	// Example: fmt.Println("Sunday:", Sunday) prints "Sunday: 0"

	// TODO: Print StatusOK, StatusCreated, and StatusAccepted with their values

	fmt.Println(Sunday, Monday, Tuesday)
	fmt.Println(StatusOK, StatusCreated, StatusAccepted)
}
