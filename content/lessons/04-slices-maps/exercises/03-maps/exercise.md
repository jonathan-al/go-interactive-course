# Exercise 03: Maps

## Objective

Learn how to create a map, add entries, iterate over it with sorted keys, and delete entries.

## Instructions

1. Create a `map[string]int` using `make`.
2. Add entries: `"Alice": 30`, `"Bob": 25`, `"Charlie": 35`.
3. Iterate over the map with **sorted keys** and print each `key: value`.
4. Delete the `"Bob"` entry.
5. Print the length of the map after deletion.

## Hint

Use `sort.Strings()` from the `sort` package to sort the map keys before iterating.

## Expected Output

```
Alice: 30
Bob: 25
Charlie: 35
Length after delete: 2
```
