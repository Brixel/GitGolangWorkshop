package ggworkshop

import (
	"bufio"
	"io"
	"os"
)

// Maze represents a maze.
type Maze struct {
	width  int
	height int
	Exits  []*Tile
	Map    [][]*Tile
}

// FromReader creates a Maze by reading the data from a io.Reader.
func FromReader(r io.Reader) *Maze {
	s := bufio.NewScanner(r)

	m := &Maze{Map: make([][]*Tile, 0)}

	for y := 0; s.Scan(); y++ {
		m.Map = append(m.Map, make([]*Tile, 0))
		bs := s.Bytes()

		for x, b := range bs {
			c := &Tile{}
			if IsPath(b) {
				c.Type = Road
			} else {
				c.Type = Wall
			}

			c.X = x
			c.Y = y
			m.Map[y] = append(m.Map[y], c)
		}
	}

	m.width = len(m.Map[0])
	m.height = len(m.Map)
	m.FindExits()

	return m
}

// FromFile creates a Maze by reading the data from a file at a location.
func FromFile(path string) (*Maze, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return FromReader(file), nil
}

// Range iterates over all the coordiantes of a Maze from x=0 to x=max
// and y=0 to y=max. If the func f returns false the iteration is stopped.
func (m *Maze) Range(f func(x, y int) bool) {
	// THIS IS THE METHOD
}

// FindExists searches for all the exits or entrances of the maze and returns
// the pointers to them in a slice.
func (m *Maze) FindExits() []*Tile {
	return nil
}

// Explore explores the maze and searches for the solution.
func (m *Maze) Explore(at *Tile, p Path, out chan<- []Path) {
}

// WaitForExporers waits for the exploring channels that are exploring
// the maze.
func (m *Maze) WaitForExplorers(explorers []chan []Path) []Path {
	return nil
}

// Look looks to a direcion (to) from a location (at) and returns what you
// are looking at (TileType) and the pointer to the Tile you are looking at.
func (m *Maze) Look(to Direction, at *Tile, p Path) (TileType, *Tile) {
	return 0, nil
}

// String draws the maze as ascii characters as a 2D map.
func (m *Maze) String() string {
	return "TODO format maze"
}
