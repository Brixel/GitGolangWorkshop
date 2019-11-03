package ggworkshop

// Direction is an int enum that represents a direction in a maze
// (North, East, South and West).
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// Directions is a slice of maze.Direction.
type Directions []Direction

var (
	// PossibleDirections is a Directions that contains all the possible
	// directions the enum can have.
	PossibleDirections Directions
)

func init() {
	PossibleDirections = make(Directions, 4, 4)
	PossibleDirections[North] = North
	PossibleDirections[East] = East
	PossibleDirections[South] = South
	PossibleDirections[West] = West
}

// Contains returns whether Directions contains a given direction.
func (ds Directions) Contains(find Direction) bool {
	for _, x := range ds {
		if x == find {
			return true
		}
	}
	return false
}

// String formats the Direction. It returns the name of the enum value.
func (d Direction) String() string {
	return "TODO: format direction"
}
