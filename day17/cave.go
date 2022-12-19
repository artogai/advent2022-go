package day17

import (
	"fmt"
	"math"
	"strings"
)

type jetDir rune

type cave struct {
	field  [][]rune
	jets   *cyclicBuff[jetDir]
	width  int
	height int
	top    int
}

func (c *cave) drop(n int) {
	for i := 0; i < n; i++ {
		switch i % 5 {
		case 0:
			c.dropShape(newHLine(c.top))
		case 1:
			c.dropShape(newCross(c.top))
		case 2:
			c.dropShape(newLShape(c.top))
		case 3:
			c.dropShape(newVLine(c.top))
		case 4:
			c.dropShape(newSquare(c.top))
		}
	}
}

func (c *cave) dropShape(s shape) {
	nextPoints := s.points()
	landed := false

	for !landed {
		switch c.jets.next() {
		case '<':
			p, moveOk := move(nextPoints, -1, 0, c.field)
			if moveOk {
				nextPoints = p
			}
		case '>':
			p, moveOk := move(nextPoints, +1, 0, c.field)
			if moveOk {
				nextPoints = p
			}
		default:
			panic("not matched")
		}

		p, moveOk := move(nextPoints, 0, +1, c.field)
		if moveOk {
			nextPoints = p
		} else {
			landed = true
		}
	}
	c.add(nextPoints)
}

func newCave(width int, height int, jets []jetDir) *cave {
	field := make([][]rune, height)
	for i := 0; i < height; i++ {
		field[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			field[i][j] = '.'
		}
	}
	return &cave{field, newCyclicBuff(jets), width, height, height}
}

func (c *cave) add(points []point) {
	minY := math.MaxInt32
	for _, p := range points {
		c.field[p.y][p.x] = '#'
		if p.y < minY {
			minY = p.y
		}
	}
	if minY < c.top {
		c.top = minY
	}
}

func (c *cave) print() {
	var str strings.Builder
	for _, row := range c.field {
		for _, cell := range row {
			str.WriteRune(cell)
		}
		str.WriteString("\n")
	}
	fmt.Println(str.String())
}
