# Maze

## Stringer implementation

When a value is printed to the console, the `Stringer` interface is used to
format the value.

A `Maze` should format to a 2D representation of the maze.

### Example

A `Maze` from a file that looks like this:

``` text
10111
10001
11101
10001
11101
```

Should return the string

``` text
x xxx
x   x
xxx x
x   x
xxx x
```

## Range(f func(x, y int) bool)

`Range` should range over all the tiles of the maze. Moving from `x=0` to `x=max`
and `y=0` to `y=max`.

## FindExits()

`FindExits` should search for all the exits in the `Maze` and set the `Exits` field
of the `Maze`.

## Look(to Direction, at *Tile, p Path) (TileType, *Tile)

Look should look to a direcion (`to`) from a location (`at`) and return what you
are looking at (`TileType`) and the pointer to the `Tile` you are looking at.
