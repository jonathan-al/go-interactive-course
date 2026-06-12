# Lesson 04: Arrays, Slices, Maps, and Iteration

## Arrays

An array is a fixed-size sequence of elements of the same type. The size is part of the type, so `[5]int` and `[10]int` are different types.

```go
var numbers [5]int
numbers[0] = 10
numbers[1] = 20

// Array literal
primes := [5]int{2, 3, 5, 7, 11}

// Let the compiler count
values := [...]int{1, 2, 3} // type is [3]int
```

Use `len()` to get the number of elements:

```go
fmt.Println(len(primes)) // 5
```

Arrays in Go are **values**, not pointers. Assigning one array to another copies all elements.

## Slices

A slice is a dynamically-sized, flexible view over an array. Slices are far more common than arrays in Go programs.

```go
// Slice literal (no size specified)
s := []int{1, 2, 3}

// Using make: make([]T, length, capacity)
s = make([]int, 5)     // len=5, cap=5
s = make([]int, 0, 10) // len=0, cap=10
```

### Appending

Use `append` to add elements. If the underlying array is too small, a new one is allocated.

```go
s := []int{1, 2, 3}
s = append(s, 4)
s = append(s, 5, 6) // append multiple elements
```

### Slicing

You can create sub-slices with the `[low:high]` syntax. Low is inclusive, high is exclusive.

```go
s := []int{1, 2, 3, 4, 5}
sub := s[1:3]  // [2, 3]
sub = s[:3]    // [1, 2, 3]
sub = s[2:]    // [3, 4, 5]
```

Slices share the same underlying array when created from slicing, so modifying one can affect the other.

## Maps

A map is an unordered collection of key-value pairs. Keys must be comparable; values can be any type.

```go
// Map literal
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}

// Using make
scores := make(map[string]int)
scores["Alice"] = 95
```

### Common operations

```go
// Get a value
age := ages["Alice"] // 30

// Check if a key exists
age, ok := ages["Charlie"]
if !ok {
    fmt.Println("Charlie not found")
}

// Delete a key
delete(ages, "Bob")

// Get the number of entries
fmt.Println(len(ages))
```

The zero value of a map is `nil`. You must initialize a map (with `make` or a literal) before writing to it.

## Iteration with range

The `range` keyword lets you iterate over slices, maps, strings, and channels.

### Iterating over a slice

```go
fruits := []string{"apple", "banana", "cherry"}
for i, v := range fruits {
    fmt.Printf("%d: %s\n", i, v)
}
```

`range` returns two values per iteration: the **index** and a **copy of the value**.

### Iterating over a map

```go
ages := map[string]int{"Alice": 30, "Bob": 25}
for name, age := range ages {
    fmt.Printf("%s is %d\n", name, age)
}
```

Map iteration order is **randomized** in Go. If you need sorted output, collect the keys, sort them, then iterate.

### Skipping values

Use `_` to ignore values you don't need:

```go
for _, v := range fruits { // skip index
    fmt.Println(v)
}

for i := range fruits { // skip value
    fmt.Println(i)
}
```

## Summary

| Concept | Key Points |
|---------|-----------|
| Arrays | Fixed size, value type, rarely used directly |
| Slices | Dynamic size, reference to array, use `append` |
| Maps | Key-value pairs, unordered, use `make` or literals |
| Range | Iterate with `for i, v := range collection` |
