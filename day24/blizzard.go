package day24

import (
	"advent2022/file"
	"math"

	"golang.org/x/exp/slices"
)

func shortestPath(from coordinate, to coordinate, c cave) int {
	caves := genCaves(c)
	path := math.MaxInt
	cache := make([][][]int, len(caves))

	for k := range caves {
		cache[k] = make([][]int, caves[k].rows)
		for i := 0; i < caves[0].rows; i++ {
			cache[k][i] = make([]int, caves[k].columns)
			for j := 0; j < caves[0].columns; j++ {
				cache[k][i][j] = math.MaxInt
			}
		}
	}

	var rec func(coordinate, int)
	rec = func(pos coordinate, steps int) {
		if pos == to {
			if steps < path {
				path = steps
			}
			return
		}
		if steps >= path {
			return
		}
		if cache[steps%len(caves)][pos.row][pos.column] <= steps {
			return
		}
		cache[steps%len(caves)][pos.row][pos.column] = steps

		c := caves[(steps+1)%len(caves)]

		for _, d := range []coordinate{{0, 0}, {-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			n := coordinate{pos.row + d.row, pos.column + d.column}
			if n != to && (n.row < 0 || n.row > c.rows-1 || n.column < 0 || n.column > c.columns-1) {
				continue
			}
			if _, ok := c.field[n.row][n.column].(open); !ok {
				continue
			}
			rec(n, steps+1)
		}
	}

	rec(from, 0)
	return path
}

type tile interface{ isTile() bool }
type blizzards struct{ n, s, w, e int }
type open struct{}
type wall struct{}

type coordinate struct {
	row    int
	column int
}

var (
	wallTile       = wall{}
	openTile       = open{}
	emptyBlizzards = blizzards{}
)

type cave struct {
	field         [][]tile
	rows, columns int
}

func newCave(rows, columns int) cave {
	field := make([][]tile, rows)
	for i := range field {
		field[i] = make([]tile, columns)
		for j := range field[i] {
			field[i][j] = openTile
		}
	}
	return cave{field, rows, columns}
}

func genCaves(c cave) []cave {
	caves := []cave{c}
	for {
		c = next(c)
		if c.eq(caves[0]) {
			break
		}
		caves = append(caves, c)
	}
	return caves
}

func next(c cave) cave {
	cnext := newCave(c.rows, c.columns)
	for i := range c.field {
		for j := range c.field[i] {
			switch tl := c.field[i][j].(type) {
			case wall:
				cnext.field[i][j] = wallTile
			case blizzards:
				move(tl, i, j, cnext)
			}
		}
	}
	return cnext
}

func move(b blizzards, i, j int, c cave) {
	update := func(i, j int, c cave, f func(blizzards) blizzards) {
		switch tl := c.field[i][j].(type) {
		case open:
			c.field[i][j] = f(emptyBlizzards)
		case blizzards:
			c.field[i][j] = f(tl)
		case wall:
			panic("move: can't be a wall")
		}
	}

	for k := 0; k < b.n; k++ {
		nextI := i
		nextJ := j
		if i > 1 {
			nextI = i - 1
		} else {
			nextI = c.rows - 2
		}
		update(nextI, nextJ, c, func(b blizzards) blizzards {
			return blizzards{b.n + 1, b.s, b.w, b.e}
		})
	}
	for k := 0; k < b.s; k++ {
		nextI := i
		nextJ := j
		if i < c.rows-2 {
			nextI = i + 1
		} else {
			nextI = 1
		}
		update(nextI, nextJ, c, func(b blizzards) blizzards {
			return blizzards{b.n, b.s + 1, b.w, b.e}
		})
	}
	for k := 0; k < b.w; k++ {
		nextI := i
		nextJ := j
		if j > 1 {
			nextJ = j - 1
		} else {
			nextJ = c.columns - 2
		}
		update(nextI, nextJ, c, func(b blizzards) blizzards {
			return blizzards{b.n, b.s, b.w + 1, b.e}
		})
	}
	for k := 0; k < b.e; k++ {
		nextI := i
		nextJ := j
		if j < c.columns-2 {
			nextJ = j + 1
		} else {
			nextJ = 1
		}
		update(nextI, nextJ, c, func(b blizzards) blizzards {
			return blizzards{b.n, b.s, b.w, b.e + 1}
		})
	}
}

func parse(filename string) cave {
	lines := file.ReadFileLines(filename)
	field := make([][]tile, len(lines))
	for i, line := range lines {
		field[i] = make([]tile, len(line))
		for j, char := range line {
			switch char {
			case '.':
				field[i][j] = openTile
			case '#':
				field[i][j] = wallTile
			case '<':
				field[i][j] = blizzards{w: 1}
			case '>':
				field[i][j] = blizzards{e: 1}
			case '^':
				field[i][j] = blizzards{n: 1}
			case 'v':
				field[i][j] = blizzards{s: 1}
			}
		}
	}
	return cave{field, len(lines), len(lines[0])}
}

func (c cave) eq(o cave) bool {
	if c.rows != o.rows || c.columns != o.columns {
		return false
	}
	for i := range c.field {
		if !slices.EqualFunc(c.field[i], o.field[i], eq) {
			return false
		}
	}
	return true
}

func eq(t1 tile, t2 tile) bool {
	switch t1.(type) {
	case open:
		_, ok := t2.(open)
		return ok
	case wall:
		_, ok := t2.(wall)
		return ok
	case blizzards:
		t2, ok := t2.(blizzards)
		return ok && t1 == t2
	}
	return false
}

func (b blizzards) isTile() bool { return true }
func (o open) isTile() bool      { return true }
func (w wall) isTile() bool      { return true }
