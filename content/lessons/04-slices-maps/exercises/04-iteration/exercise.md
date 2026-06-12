# Exercise 04: Iteration with Range

## Objective

Learn how to use `range` to iterate over slices and maps.

## Instructions

1. Create a slice of strings: `"apple", "banana", "cherry"`.
2. Use `range` to iterate and print each element as `index: value`.
3. Create a `map[int]string`: `1: "one", 2: "two", 3: "three"`.
4. Iterate over the map with **sorted keys** and print each entry as `key = value`.

## Hint

Collect map keys into a slice, sort them with `sort.Ints()`, then iterate.

## Expected Output

```
0: apple
1: banana
2: cherry
1 = one
2 = two
3 = three
```
