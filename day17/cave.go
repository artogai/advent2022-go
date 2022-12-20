package day17

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type jetDir bool

type cave struct {
	field    *[][]bool
	jets     *cyclicBuff[jetDir]
	width    int
	height   int
	top      int
	fullLine int
	skipped  int
}

func (c *cave) dropSkipping(n int) {
	beforeCycle := 347
	cycleIter := 1745
	c.drop(0, beforeCycle)
	skip := (n - beforeCycle) / cycleIter
	c.skipped += skip * 2753
	conitnue := n - beforeCycle - (skip * cycleIter)

	c.drop(beforeCycle+skip*cycleIter, beforeCycle+skip*cycleIter+conitnue)
}

func (c *cave) drop(from int, until int) {
	hline := new(hline)
	cross := new(cross)
	lshape := new(lshape)
	vline := new(vline)
	square := new(square)
	var currShape shape

	for i := from; i < until; i++ {
		switch i % 5 {
		case 0:
			currShape = hline
		case 1:
			currShape = cross
		case 2:
			currShape = lshape
		case 3:
			currShape = vline
		case 4:
			currShape = square
		default:
			panic("unknown shape")
		}
		currShape.reset(c.top)
		c.dropShape(currShape)
	}
}

func (c *cave) dropShape(s shape) {
	for {
		if c.jets.next() {
			_ = move(s, 1, 0, c)
		} else {
			_ = move(s, -1, 0, c)
		}
		if ok := move(s, 0, 1, c); !ok {
			break
		}
	}
	c.add(s)
}

func (c *cave) add(s shape) {
	prevTop := c.top
	for _, p := range s.points() {
		(*c.field)[p.y][p.x] = true
		if c.top > p.y {
			c.top = p.y
		}
	}
	topDiff := prevTop - c.top

	if topDiff > 0 {
		c.clearTop(topDiff)
	}
	if fullLine, exists := s.findFullLine(c); exists && fullLine < c.fullLine {
		c.fullLine = fullLine
	}
}

func (c *cave) clearTop(topDiff int) {
	for i := 0; i < topDiff; i++ {
		for j := 0; j < c.width; j++ {
			(*c.field)[c.top-clearAboveTop+i][j] = false
		}
	}
}

func (c *cave) optimize() {
	if c.fullLine != c.height {
		i := c.fullLine
		// moves all lines above full line to the bottom
		for j := i - 1; j >= c.top; j-- {
			newY := c.height - i + j
			for k := 0; k < c.width; k++ {
				(*c.field)[newY][k] = (*c.field)[j][k]
			}
		}

		c.top = c.height - i + c.top
		c.skipped += c.height - i
		c.fullLine = c.height

		// clears only clearAboveTop lines above new top
		for j := c.top - 1; j > c.top-1-clearAboveTop; j-- {
			for k := 0; k < c.width; k++ {
				(*c.field)[j][k] = false
			}
		}
	} else {
		panic("full line not found")
	}
}

func (c *cave) towerHeight() int {
	return c.skipped + c.height - c.top
}

func newCave(width int, height int, jets []jetDir) *cave {
	field := make([][]bool, height)
	for i := 0; i < height; i++ {
		field[i] = make([]bool, width)
	}
	return &cave{
		field:    &field,
		jets:     newCyclicBuff(jets),
		width:    width,
		height:   height,
		top:      height,
		fullLine: height,
		skipped:  0,
	}
}

func (c *cave) print() {
	fmt.Println(c.toString())
}

func (c *cave) write(filename string) {
	ioutil.WriteFile(filename, []byte(c.toString()), 0644)
}

func (c *cave) toString() string {
	var str strings.Builder
	for _, row := range *c.field {
		for _, cell := range row {
			if cell {
				str.WriteRune('#')
			} else {
				str.WriteRune('.')
			}
		}
		str.WriteString("\n")
	}
	return str.String()
}
