package day14

import (
	"math"
	"strconv"
	"strings"
)

type cave [][]rune

func newCave(dim int) cave {
	c := make(cave, dim)
	for i := 0; i < dim; i++ {
		c[i] = make([]rune, dim)
		for j := 0; j < dim; j++ {
			c[i][j] = '.'
		}
	}
	return c
}

func (c cave) dropSand(x int, y int) bool {
	switch c[x][y] {
	case '.':
		// source is not blocked
		if x+1 < len(c) {
			// try fall down
			switch c[x+1][y] {
			case '.':
				// falling down
				return c.dropSand(x+1, y)
			case '#', 'o':
				if y-1 < 0 {
					// falling into abyss
					return false
				} else {
					// try fall down-left
					switch c[x+1][y-1] {
					case '.':
						// falling down-left
						return c.dropSand(x+1, y-1)
					case '#', 'o':
						if y+1 >= len(c) {
							// falling into abyss
							return false
						} else {
							// try fall down-right
							switch c[x+1][y+1] {
							case '.':
								// falling down-right
								return c.dropSand(x+1, y+1)
							case '#', 'o':
								// sand stops
								c[x][y] = 'o'
								return true
							}
						}
					}
				}
			}
		} else {
			return false
		}
	case '#', 'o':
		// source is blocked
		return false
	}
	panic("unreachable")
}

func (c cave) countSand() int {
	count := 0
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c); j++ {
			if c[i][j] == 'o' {
				count += 1
			}
		}
	}
	return count
}

func (c cave) show() string {
	var str strings.Builder
	for _, row := range c {
		for _, cell := range row {
			str.WriteRune(cell)
		}
		str.WriteString("\n")
	}
	return str.String()
}

func read(dim int, addFloor bool, strs ...string) cave {
	c := newCave(dim)
	maxX := 0
	for _, str := range strs {
		prevX := -1
		prevY := -1
		pointsStr := strings.Split(str, " -> ")
		for _, pointStr := range pointsStr {
			pointsArr := strings.Split(pointStr, ",")
			y, err := strconv.Atoi(pointsArr[0])
			if err != nil {
				panic(err)
			}
			x, err := strconv.Atoi(pointsArr[1])
			if err != nil {
				panic(err)
			}
			if x > maxX {
				maxX = x
			}
			if prevX != -1 {
				drawLine(prevX, prevY, x, y, c)
			}
			prevX = x
			prevY = y
		}
	}

	if addFloor {
		for i := 0; i < len(c); i++ {
			c[maxX+2][i] = '#'
		}
	}

	return c
}

func drawLine(prevX, prevY, x, y int, c cave) {
	if prevX == x {
		from := int(math.Min(float64(prevY), float64(y)))
		to := int(math.Max(float64(prevY), float64(y)))
		for i := from; i <= to; i++ {
			c[x][i] = '#'
		}
	} else if prevY == y {
		from := int(math.Min(float64(prevX), float64(x)))
		to := int(math.Max(float64(prevX), float64(x)))
		for i := from; i <= to; i++ {
			c[i][y] = '#'
		}
	} else {
		panic("not a line")
	}
}
