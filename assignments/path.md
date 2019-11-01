# Path

## Stringer implementation

When a value is printed to the console, the `Stringer` interface is used to
format the value.

A `Path` should format to the 2D representation of the path.

### Example

``` golang
p := Path {
    Coordinate{X: 1, Y:0}
    Coordinate{X: 1, Y:1}
    Coordinate{X: 2, Y:1}
    Coordinate{X: 2, Y:2}
    Coordinate{X: 2, Y:3}
}
s := p.String()
```

s should be:

``` golang
x xx
x  x
xx x
xx x
```

## `Last() *Coordinate`

`Last` should return the last element of the `Path`.

## `Contains(x *Coordinate)`

`Contains` should return whether a `Path` contains a given `Coordinate`.

## IsPath(b byte) bool

`IsPath` returns whether a byte representas a wall or path segement in a maze.
The different bytes that are a path are: `' '`, `0` and `byte(0)`.
