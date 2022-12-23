package day22

import (
	"advent2022/matrix"
	"fmt"
	"strings"
)

type board interface {
	step(p *Position)
	StartPosition() Position
}

type flatBoard [][]tile
type cubeBoard struct {
	field       cube[tile]
	coordinates cube[coordinate]
}

type tile int
type coordinate struct{ row, column int }

const (
	open  tile = iota
	wall  tile = iota
	empty tile = iota
)

func ParseFlatBoard(board string, rows, columns int) *flatBoard {
	b := make(flatBoard, rows)
	rws := strings.Split(board, "\n")[0:rows]
	for i, s := range rws {
		b[i] = parseRow(s, columns)
	}
	return &b
}

func NewCubeBoard(fb *flatBoard) *cubeBoard {
	edge0, coords0 := cutEdge(0, 50, fb)
	edge1, coords1 := cutEdge(0, 100, fb)
	edge5, coords5 := cutEdge(50, 50, fb)
	edge2, coords2 := cutEdge(100, 50, fb)
	edge3, coords3 := cutEdge(100, 0, fb)
	edge4, coords4 := cutEdge(150, 0, fb)

	edge2.RotateClockwise()
	edge2.RotateClockwise()
	coords2.RotateClockwise()
	coords2.RotateClockwise()

	edge3.RotateClockwise()
	edge3.RotateClockwise()
	coords3.RotateClockwise()
	coords3.RotateClockwise()

	edge4.RotateAntiClockwise()
	coords4.RotateAntiClockwise()

	return &cubeBoard{
		field: cube[tile]{
			front:  edge0,
			right:  edge1,
			back:   edge2,
			left:   edge3,
			top:    edge4,
			bottom: edge5,
		},
		coordinates: cube[coordinate]{
			front:  coords0,
			right:  coords1,
			back:   coords2,
			left:   coords3,
			top:    coords4,
			bottom: coords5,
		},
	}
}

func (fb *flatBoard) StartPosition() Position {
	for row := 0; row < len(*fb); row++ {
		for column := 0; column < len((*fb)[row]); column++ {
			if (*fb)[row][column] == open {
				return Position{row, column, east}
			}
		}
	}
	panic("StartPosition: no start position found on flat board")
}

func (cb *cubeBoard) StartPosition() Position {
	for row := 0; row < cb.field.front.Size(); row++ {
		for column := 0; column < cb.field.front.Size(); column++ {
			if cb.field.front[row][column] == open {
				return Position{row, column, east}
			}
		}
	}
	panic("StartPosition: no start position found on cube board")
}

func (cb *cubeBoard) OriginalPosition(p Position) Position {
	c := cb.coordinates.front[p.row][p.column]
	return Position{c.row, c.column, p.facing}
}

func cutEdge(row, column int, fb *flatBoard) (matrix.Square[tile], matrix.Square[coordinate]) {
	edge := make([][]tile, 50)
	coords := make([][]coordinate, 50)
	for i := range edge {
		edge[i] = make([]tile, 50)
		coords[i] = make([]coordinate, 50)
	}

	for r := row; r < row+50; r++ {
		for c := column; c < column+50; c++ {
			coords[r-row][c-column] = coordinate{row: r, column: c}
			edge[r-row][c-column] = (*fb)[r][c]
		}
	}
	return matrix.Square[tile](edge), matrix.Square[coordinate](coords)
}

func parseRow(s string, columns int) []tile {
	row := make([]tile, columns)
	for i := range row {
		row[i] = empty
	}

	for i, c := range s {
		switch c {
		case '.':
			row[i] = open
		case '#':
			row[i] = wall
		case ' ':
			row[i] = empty
		default:
			panic(fmt.Sprint("parseRow: unknown tile ", c))
		}
	}
	return row
}

func (cb *cubeBoard) step(p *Position) {
	nextRow := p.row
	nextColumn := p.column
	rotation := -1

	switch p.facing {
	case north:
		if p.row-1 < 0 {
			cb.rotateDown()
			rotation = 0
			nextRow = cb.field.edgeSize() - 1
		} else {
			nextRow = p.row - 1
		}
	case south:
		if p.row+1 >= cb.field.edgeSize() {
			cb.rotateUp()
			rotation = 1
			nextRow = 0
		} else {
			nextRow = p.row + 1
		}
	case east:
		if p.column+1 >= cb.field.edgeSize() {
			cb.rotateLeft()
			rotation = 2
			nextColumn = 0
		} else {
			nextColumn = p.column + 1
		}
	case west:
		if p.column-1 < 0 {
			cb.rotateRight()
			rotation = 3
			nextColumn = cb.field.edgeSize() - 1
		} else {
			nextColumn = p.column - 1
		}
	}

	switch cb.field.front[nextRow][nextColumn] {
	case open:
		p.row = nextRow
		p.column = nextColumn
		return
	case wall:
		switch rotation {
		case 0:
			cb.rotateUp()
		case 1:
			cb.rotateDown()
		case 2:
			cb.rotateRight()
		case 3:
			cb.rotateLeft()
		}
		return
	case empty:
		panic("step: no board tile found")
	}
}

func (fb *flatBoard) step(p *Position) {
	nextRow := p.row
	nextColumn := p.column

	switch p.facing {
	case north:
		if p.row-1 < 0 || (*fb)[p.row-1][p.column] == empty {
			nextRow = findWrapIndexNorth(p.column, fb)
		} else {
			nextRow = p.row - 1
		}
	case south:
		if p.row+1 >= len(*fb) || (*fb)[p.row+1][p.column] == empty {
			nextRow = findWrapIndexSouth(p.column, fb)
		} else {
			nextRow = p.row + 1
		}
	case east:
		if p.column+1 >= len((*fb)[p.row]) || (*fb)[p.row][p.column+1] == empty {
			nextColumn = findWrapIndexEast(p.row, fb)
		} else {
			nextColumn = p.column + 1
		}
	case west:
		if p.column-1 < 0 || (*fb)[p.row][p.column-1] == empty {
			nextColumn = findWrapIndexWest(p.row, -1, fb)
		} else {
			nextColumn = p.column - 1
		}
	}

	switch (*fb)[nextRow][nextColumn] {
	case open:
		p.row = nextRow
		p.column = nextColumn
		return
	case wall:
		return
	case empty:
		panic("step: empty tile after wrapping")
	}
}

func findWrapIndexNorth(column int, fb *flatBoard) int {
	for i := len(*fb) - 1; i >= 0; i-- {
		if (*fb)[i][column] != empty {
			return i
		}
	}
	panic("findWrapIndexNorth: no board tile found")
}

func findWrapIndexSouth(column int, fb *flatBoard) int {
	for i := 0; i < len(*fb); i++ {
		if (*fb)[i][column] != empty {
			return i
		}
	}
	panic("findWrapIndexSouth: no board tile found")
}

func findWrapIndexEast(row int, fb *flatBoard) int {
	for i := 0; i < len((*fb)[row]); i++ {
		if (*fb)[row][i] != empty {
			return i
		}
	}
	panic("findWrapIndexEast: no board tile found")
}

func findWrapIndexWest(row, column int, fb *flatBoard) int {
	for i := len((*fb)[row]) - 1; i >= 0; i-- {
		if (*fb)[row][i] != empty {
			return i
		}
	}
	panic("findWrapIndexWest: no board tile found")
}

func (cb *cubeBoard) rotateDown() {
	cb.field.rotateDown()
	cb.coordinates.rotateDown()
}

func (cb *cubeBoard) rotateUp() {
	cb.field.rotateUp()
	cb.coordinates.rotateUp()
}

func (cb *cubeBoard) rotateLeft() {
	cb.field.rotateLeft()
	cb.coordinates.rotateLeft()
}

func (cb *cubeBoard) rotateRight() {
	cb.field.rotateRight()
	cb.coordinates.rotateRight()
}
