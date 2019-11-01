# Coordinate

## Stringer implementation

When a value is printed to the console, the `Stringer` interface is used to
format the value.

A `Coordinate` should format to `($X, $Y)`.

### Example

``` golang
c := Coordinate{X: 1, Y: 4}
s := c.String()
```

`s` should be:

``` golang
"(1, 4)"
```
