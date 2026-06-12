package main

import (
	"fmt"
	"sort"
)

func main() {
	// TODO: Create a map[string]int using make
	ages := make(map[string]int)

	// TODO: Add entries - "Alice": 30, "Bob": 25, "Charlie": 35
	ages["Alice"] = 0

	// TODO: Iterate with sorted keys and print each entry
	keys := make([]string, 0, len(ages))
	for k := range ages {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, ages[k])
	}

	// TODO: Delete "Bob" and print the length
	delete(ages, "Bob")
	fmt.Println("Length after delete:", len(ages))
}
