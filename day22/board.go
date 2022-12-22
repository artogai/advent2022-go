package day22

import "fmt"

type tile int

const (
	open tile = iota
	wall
	empty
)

type direction int

const (
	north direction = iota
	east  direction = iota
	south direction = iota
	west  direction = iota
)

type board [][]tile

type Position struct {
	row, column int
	facing      direction
}

func ParseBoard(strs []string, rows, columns int) board {
	b := make(board, rows)
	for i, s := range strs {
		b[i] = parseRow(s, columns)
	}
	return b
}

func (p *Position) MoveMany(ms []move, b board) {
	for _, m := range ms {
		p.move(m, b)
	}
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

func (p *Position) move(m move, b board) {
	switch m := m.(type) {
	case rotate:
		p.rotate(m)
	case walk:
		p.walk(int(m), b)
	default:
		panic(fmt.Sprint("move: unknown move ", m))
	}
}

func (p *Position) rotate(r rotate) {
	switch r {
	case L:
		p.facing = (p.facing + 3) % 4
	case R:
		p.facing = (p.facing + 1) % 4
	}
}

func (p *Position) walk(distance int, b board) {
	for i := 0; i < distance; i++ {
		p.step(b)
	}
}

func (p *Position) step(b board) {
	nextRow := p.row
	nextColumn := p.column

	switch p.facing {
	case north:
		if p.row-1 < 0 || b[p.row-1][p.column] == empty {
			nextRow = findWrapIndexNorth(p.column, b)
		} else {
			nextRow = p.row - 1
		}
	case south:
		if p.row+1 >= len(b) || b[p.row+1][p.column] == empty {
			nextRow = findWrapIndexSouth(p.column, b)
		} else {
			nextRow = p.row + 1
		}
	case east:
		if p.column+1 >= len(b[p.row]) || b[p.row][p.column+1] == empty {
			nextColumn = findWrapIndexEast(p.row, b)
		} else {
			nextColumn = p.column + 1
		}
	case west:
		if p.column-1 < 0 || b[p.row][p.column-1] == empty {
			nextColumn = findWrapIndexWest(p.row, -1, b)
		} else {
			nextColumn = p.column - 1
		}
	}

	switch b[nextRow][nextColumn] {
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

func findWrapIndexNorth(column int, b board) int {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i][column] != empty {
			return i
		}
	}
	panic("findWrapIndexNorth: no board tile found")
}

func findWrapIndexSouth(column int, b board) int {
	for i := 0; i < len(b); i++ {
		if b[i][column] != empty {
			return i
		}
	}
	panic("findWrapIndexSouth: no board tile found")
}

func findWrapIndexEast(row int, b board) int {
	for i := 0; i < len(b[row]); i++ {
		if b[row][i] != empty {
			return i
		}
	}
	panic("findWrapIndexEast: no board tile found")
}

func findWrapIndexWest(row, column int, b board) int {
	for i := len(b[row]) - 1; i >= 0; i-- {
		if b[row][i] != empty {
			return i
		}
	}
	panic("findWrapIndexWest: no board tile found")
}
