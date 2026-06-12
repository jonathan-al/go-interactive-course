# Exercise 4: The fmt Package

## Objective

Practice using `fmt.Sprintf`, `fmt.Printf`, and format verbs to produce formatted output.

## Instructions

Open `starter.go`. Several variables are declared and printed with basic `fmt.Println` calls. Your task is to **replace the print statements** with properly formatted output using `fmt.Sprintf`, `fmt.Printf`, and the correct format verbs.

Produce this exact output:

```
Name: Bob, Age: 25
Score: 100
Pi: 3.14
Active: true
Type of name: string
```

## Requirements

1. Use `fmt.Sprintf` to build the first line (`"Name: Bob, Age: 25"`) and print it with `fmt.Println`.
2. Use `fmt.Printf` with `%d` to print the score.
3. Use `fmt.Printf` with `%.2f` to print pi with 2 decimal places.
4. Use `fmt.Printf` with `%t` to print the boolean.
5. Use `fmt.Printf` with `%T` to print the type of `name`.

## Expected Output

```
Name: Bob, Age: 25
Score: 100
Pi: 3.14
Active: true
Type of name: string
```

## Hints

- `fmt.Sprintf` returns a string — use `fmt.Println` to print it.
- `%d` formats integers, `%.2f` formats floats with 2 decimal places.
- `%t` formats booleans, `%T` prints the type of a value.
- Remember `\n` when using `fmt.Printf`.
