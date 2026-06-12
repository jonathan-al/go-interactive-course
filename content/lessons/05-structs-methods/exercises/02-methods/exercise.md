# Exercise 02: Methods

## Objective

Learn how to define methods on a struct, including value receivers and pointer receivers.

## Instructions

1. Define a `Rectangle` struct with `Width` and `Height` fields of type `float64`.
2. Add an `Area()` method with a **value receiver** that returns the area (`Width * Height`).
3. Add a `Scale(factor float64)` method with a **pointer receiver** that multiplies both `Width` and `Height` by the given factor.
4. Create a `Rectangle` with `Width: 5` and `Height: 3`.
5. Print the dimensions and area.
6. Scale the rectangle by a factor of `2`.
7. Print the new dimensions and area.

## Expected Output

```
Width: 5.00, Height: 3.00
Area: 15.00
After scaling by 2.00:
Width: 10.00, Height: 6.00
Area: 60.00
```
