# Exercise 3: Types

## Objective

Explore Go's basic types and learn how to inspect the type of a value.

## Instructions

Open `starter.go`. Four variables are declared with different types and printed using `fmt.Println`. Your task is to **replace the `fmt.Println` calls with `fmt.Printf`** to display each variable's type and value.

Use the `%T` format verb for type and `%v` for value. Each line should follow this format:

```
a: type=int, value=42
b: type=float64, value=3.14
c: type=string, value=Hello
d: type=bool, value=true
```

## Expected Output

```
a: type=int, value=42
b: type=float64, value=3.14
c: type=string, value=Hello
d: type=bool, value=true
```

## Hints

- `fmt.Printf` does NOT add a newline — include `\n` at the end of your format string.
- Example: `fmt.Printf("a: type=%T, value=%v\n", a, a)`
- `%T` prints the type of a value, `%v` prints its default representation.
