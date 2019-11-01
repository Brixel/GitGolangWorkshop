# TileType

## Stringer implementation

When a value is printed to the console, the `Stringer` interface is used to
format the value.

A `TileType` should format to `$Name` where `$Name` is the enum name of the
`TileType`.

### Example

``` golang
t := Road
s := t.String()
```

`s` should be:

``` golang
"Road"
```
