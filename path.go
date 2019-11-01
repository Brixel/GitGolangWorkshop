package ggworkshop

var (
	pathChars [3]byte
)

func init() {
	pathChars[0] = ' '
	pathChars[1] = '0'
	pathChars[2] = byte(0)
}

// IsPath indicates whether a byte represents a wall or path segment
// (a path can be ' ', '0' or the byte value 0).
func IsPath(b byte) bool {
	return false
}

// Path is a slice of Coordinate that represents a route in a maze.
type Path []*Coordinate

// Last returns the pointer to the last Coordinate of the Path.
func (p Path) Last() *Coordinate {
	return nil
}

// Contains checks whether a path goes through a given coordinate.
func (p Path) Contains(x *Coordinate) bool {
	// THIS IS THE METHODS
	return false
}

// String formats a path to a string. It visualizes the path in a 2D grid.
func (p Path) String() string {
	return "TODO: format path"
}
