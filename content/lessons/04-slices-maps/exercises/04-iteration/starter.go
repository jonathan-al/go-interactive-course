package main

import (
	"fmt"
	"sort"
)

func main() {
	// TODO: Create a slice of strings: "apple", "banana", "cherry"
	fruits := []string{"apple"}

	// TODO: Use range to iterate and print "index: value"
	for i, v := range fruits {
		fmt.Printf("%d: %s\n", i, v)
	}

	// TODO: Create a map[int]string: 1: "one", 2: "two", 3: "three"
	numbers := map[int]string{1: "one"}

	// TODO: Iterate over the map with sorted keys and print "key = value"
	keys := make([]int, 0, len(numbers))
	for k := range numbers {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%d = %s\n", k, numbers[k])
	}
}
