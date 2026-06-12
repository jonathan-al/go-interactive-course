# Exercise: Switch - Day of the Week

## Objective

Use a `switch` statement to print the name of a day given its number (1 = Monday, 7 = Sunday).

## Instructions

1. Open `starter.go`.
2. The variable `dayNumber` is already set to `3`.
3. Replace the `TODO` comment with a `switch` statement on `dayNumber` that prints the corresponding day name:
   - 1 → Monday
   - 2 → Tuesday
   - 3 → Wednesday
   - 4 → Thursday
   - 5 → Friday
   - 6 → Saturday
   - 7 → Sunday
4. Include a `default` case that prints `"Invalid day number"`.

## Hints

- Each `case` should call `fmt.Println` with the day name.
- Remember: Go switch cases do not fall through by default.

## Expected Output

```
Wednesday
```
