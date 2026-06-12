# Exercise: For Loops - Counting and Ranging

## Objective

Practice Go's `for` loop in its classic form and with `range`.

## Instructions

1. Open `starter.go`.
2. **Part 1:** Replace the first `TODO` with a classic `for` loop that prints the numbers 1 through 5, one per line.
3. **Part 2:** Replace the second `TODO` with a `for ... range` loop over the `fruits` slice that prints each fruit with its index in the format `"index: value"`.

## Hints

- Classic loop: `for i := 1; i <= 5; i++ { ... }`
- Range loop: `for i, v := range fruits { ... }`
- Use `fmt.Printf("%d: %s\n", i, v)` for the formatted output.

## Expected Output

```
1
2
3
4
5
0: apple
1: banana
2: cherry
```
