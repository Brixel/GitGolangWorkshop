package ggworkshop

// TileType is an enum int that represents the the type of a tyle
// (Wall, Road, Exit, Duplicate).
type TileType int

// The value of a TileType
const (
	Wall TileType = iota
	Road
	Exit
	Duplicate
)

// Tile represents a location in a maze.
type Tile struct {
	Coordinate
	Type TileType
}

// String formats a TyleType as a string. It returns the enum name.
func (t TileType) String() string {
	switch t {
	case Wall:
		return "Wall"

	case Road:
		return "Road"

	case Exit:
		return "Exit"

	case Duplicate:
		return "Duplicate"

	default:
		return "error,wrong input"
	}
}
