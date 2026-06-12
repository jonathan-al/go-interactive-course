# Exercise: If / Else - Number Sign Check

## Objective

Write a program that checks whether a number is **positive**, **negative**, or **zero** using `if / else if / else`.

## Instructions

1. Open `starter.go`.
2. The variable `num` is already defined with a value of `-5`.
3. Replace the `TODO` comment with an `if / else if / else` block that:
   - Prints `"The number -5 is positive"` if `num` is greater than zero.
   - Prints `"The number -5 is negative"` if `num` is less than zero.
   - Prints `"The number is zero"` if `num` equals zero.
4. Use `fmt.Printf` so the number is included in the output for the positive/negative cases.

## Hints

- Use `num > 0`, `num < 0`, and `num == 0` as conditions.
- `fmt.Printf("The number %d is positive\n", num)` uses `%d` to format an integer.

## Expected Output

```
The number -5 is negative
```
