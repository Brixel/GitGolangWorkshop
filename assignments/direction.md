# Direction

## Stringer implementation

When a value is printed to the console, the `Stringer` interface is used to
format the value.

A `Direction` should format to `$Name` where `$Name` is the enum name of the
`Direction`.

### Example

``` golang
d := North
s := d.String()
```

`s` should be:

``` golang
"North"
```

## Contains()

Contains should return whether a `Direction` is present in the `Directions`. 

### Example

``` golang
ds := Directions{
    North,
    East,
}
t := ds.Contains(North)
f := ds.Contains(East)
```

`t` should be `true` and `f` should be `false`

