package ggworkshop

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func CompareMaze(m *Maze, bs [][]bool) error {
	var err error
	m.Range(func(x, y int) bool {
		switch m.Map[y][x].Type {
		case Road, Exit:
			if !bs[y][x] {
				err = errors.New(fmt.Sprintf("(%d, %d) => (%v) is not a wall", x, y, m.Map[y][x].Type))
			}
		case Wall:
			if bs[y][x] {
				err = errors.New(fmt.Sprintf("(%d, %d) => (%v) is a wall", x, y, m.Map[y][x].Type))
			}
		case Duplicate:
			err = errors.New(fmt.Sprintf("Duplicate (%d, %d)\r\n", x, y))
		}
		return false
	})

	return err
}

func PrintBools(bs [][]bool) {
	builder := strings.Builder{}

	for y := range bs {
		for x := range bs[y] {
			if bs[y][x] {
				builder.WriteString(".")
			} else {
				builder.WriteString("x")
			}
		}
		builder.WriteString("\r\n")
	}

	fmt.Println(builder.String())
}

func TestParse4x4(t *testing.T) {
	m, err := FromFile("./mazes/4x4.maze")
	if err != nil {
		t.Error(err)
		return
	}

	bs := [][]bool{
		{false, true, false, false},
		{false, true, true, false},
		{false, false, true, false},
		{false, false, true, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		PrintBools(bs)
		fmt.Println(m)
		t.Error(err)
	}
}

func TestParse5x5(t *testing.T) {
	m, err := FromFile("./mazes/5x5.maze")
	if err != nil {
		t.Error(err)
		return
	}

	bs := [][]bool{
		{false, true, false, false, false},
		{false, true, true, true, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
		{false, false, false, true, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		fmt.Println(m)
		t.Error(err)
	}
}

func TestParse6x6(t *testing.T) {
	m, err := FromFile("./mazes/6x6.maze")
	if err != nil {
		t.Error(err)
		return
	}

	bs := [][]bool{
		{false, true, false, false, false, false},
		{false, true, false, false, false, false},
		{false, true, true, true, false, false},
		{false, false, true, false, true, false},
		{false, false, true, true, true, false},
		{false, false, false, true, false, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		fmt.Println(m)
		t.Error(err)
	}
}

func TestExplore4x4(t *testing.T) {
	m, err := FromFile("./mazes/4x4.maze")
	if err != nil {
		t.Error(err)
	}
	out := make(chan []Path)

	go m.Explore(m.Map[0][1], Path{}, out)

	paths := <-out
	p := paths[1]
	if len(p) != 5 {
		fmt.Println(len(p))
		fmt.Println(p)
		t.Errorf("The length of the route should be 5, not %d", len(p))
	}
}

func TestRange(t *testing.T) {
	m, err := FromFile("./mazes/4x4.maze")
	if err != nil {
		t.Error(err)
		return
	}

	var count int
	m.Range(func(x, y int) bool {
		count++
		return true
	})

	if count != 16 {
		t.Errorf("Ranges counts to 16 not %d", count)
	}

	count = 0
	m.Range(func(x, y int) bool {
		count++
		return count != 10
	})

	if count != 10 {
		t.Errorf("The range is terminated when the count is 10 not %d", count)
	}
}

func TestFindExits(t *testing.T) {
	m, err := FromFile("./mazes/4x4.maze")
	if err != nil {
		t.Error(err)
		return
	}

	if len(m.Exits) != 2 {
		t.Errorf("There are 2 exits, not %d", len(m.Exits))
	}

	if m.Exits[0].Coordinate.X != 1 || m.Exits[0].Coordinate.Y != 0 {
		t.Errorf("The first exit has (X,Y)(1,0), not (%d, %d)", m.Exits[0].Coordinate.X, m.Exits[0].Coordinate.Y)
	}
	if m.Exits[1].Coordinate.X != 2 || m.Exits[1].Coordinate.Y != 3 {
		t.Errorf("The first exit has (X,Y)(2,3), not (%d, %d)", m.Exits[1].Coordinate.X, m.Exits[1].Coordinate.Y)
	}
}

func TestLook(t *testing.T) {
	m, err := FromFile("./mazes/4x4.maze")
	if err != nil {
		t.Error(err)
		return
	}

	tileType, tile := m.Look(South, m.Map[0][1], Path{})
	if tileType != Wall {
		t.Errorf("We are looking at a wall, not %s", tileType)
	}

	if tile.Y != 1 {
		t.Errorf("We moved south. The tile should be at Y=1 not %d", tile.Y)
	}

	tileType, tile = m.Look(South, m.Map[0][1], Path{&m.Map[1][1].Coordinate})
	if tileType != Duplicate {
		t.Errorf("We are looking at a wall, not %s", tileType)
	}
}
