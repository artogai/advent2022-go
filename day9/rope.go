package day9

import (
	"advent2022/file"
	"fmt"
	"strconv"
	"strings"
)

func CalcVisited(filename string, ropeSize int, fieldSize int) int {
	visited := make([][]bool, fieldSize)
	for i := range visited {
		visited[i] = make([]bool, fieldSize)
	}

	pos := fieldSize / 2
	dirs := readDirections(filename)
	rope := newRope(pos, pos, ropeSize)

	for _, dir := range dirs {
		rope.traverse(dir, visited, nil)
	}

	return calcVisited(visited)
}

func calcVisited(visited [][]bool) int {
	cnt := 0
	for _, row := range visited {
		for _, v := range row {
			if v {
				cnt += 1
			}
		}
	}
	return cnt
}

type direction int

const (
	up    direction = iota
	down  direction = iota
	left  direction = iota
	right direction = iota
)

type knot struct{ x, y int }
type rope struct{ knots []knot }

func newRope(x, y int, len int) rope {
	knots := make([]knot, 0, len)
	for i := 0; i < len; i++ {
		knots = append(knots, knot{x, y})
	}
	return rope{knots}
}

func (r rope) traverse(dir direction, visited [][]bool, screen [][]rune) {
	if screen != nil {
		r.clear(screen)
	}
	r.move(dir)
	r.drawTail(visited)
	if screen != nil {
		r.draw(screen)
	}
}

func (r rope) move(dir direction) {
	prevLead := r.knots[0]
	switch dir {
	case up:
		r.knots[0].x -= 1
	case down:
		r.knots[0].x += 1
	case left:
		r.knots[0].y -= 1
	case right:
		r.knots[0].y += 1
	}

	for i := 1; i < len(r.knots); i++ {
		lead := r.knots[i-1]
		temp := r.knots[i]
		r.knots[i].move(prevLead, lead)
		prevLead = temp
	}
}

func (k *knot) move(prevLead knot, lead knot) {
	/*
		H - lead
		M - lead, special case
		O - prev lead
		T - tail
	*/
	if k.x == lead.x-2 || k.x == lead.x+2 ||
		k.y == lead.y+2 || k.y == lead.y-2 {

		if k.x == prevLead.x && k.y == prevLead.y-1 {
			/*
				......
				....M.
				..TO..
				....M.
				......
			*/
			if k.x == lead.x+1 && k.y == lead.y-2 {
				k.x -= 1
				k.y += 1
			} else if k.x == lead.x-1 && k.y == lead.y-2 {
				k.x += 1
				k.y += 1
			} else {
				k.x = prevLead.x
				k.y = prevLead.y
			}
		} else if k.x == prevLead.x-1 && k.y == prevLead.y {
			/*
				......
				......
				..T...
				..O...
				.M.M..
			*/
			if k.x == lead.x-2 && k.y == lead.y+1 {
				k.x += 1
				k.y -= 1
			} else if k.x == lead.x-2 && k.y == lead.y-1 {
				k.x += 1
				k.y += 1
			} else {
				k.x = prevLead.x
				k.y = prevLead.y
			}
		} else if k.x == prevLead.x && k.y == prevLead.y+1 {
			/*
				......
				.M....
				..OT..
				.M....
				......
			*/
			if k.x == lead.x-1 && k.y == lead.y+2 {
				k.x += 1
				k.y -= 1
			} else if k.x == lead.x+1 && k.y == lead.y+2 {
				k.x -= 1
				k.y -= 1
			} else {
				k.x = prevLead.x
				k.y = prevLead.y
			}
		} else if k.x == prevLead.x+1 && k.y == prevLead.y {
			/*
				..M.M.
				...O..
				...T..
				.,....
				......
			*/
			if k.x == lead.x+2 && k.y == lead.y+1 {
				k.x -= 1
				k.y -= 1
			} else if k.x == lead.x+2 && k.y == lead.y-1 {
				k.x -= 1
				k.y += 1
			} else {
				k.x = prevLead.x
				k.y = prevLead.y
			}
		} else if k.x == prevLead.x-1 && k.y == prevLead.y+1 {
			/*
				.........
				...MHT...
				...HOH...
				...HHM...
				.........
			*/
			if k.x == lead.x && k.y == lead.y+2 {
				k.y -= 1
			} else if k.x == lead.x-2 && k.y == lead.y {
				k.x += 1
			} else {
				k.x = prevLead.x
				k.y = prevLead.y
			}
		} else if k.x == prevLead.x+1 && k.y == prevLead.y+1 {
			/*
				.........
				..HHM....
				..HOH....
				..MHT....
				.........
			*/
			if k.x == lead.x && k.y == lead.y+2 {
				k.y -= 1
			} else if k.x == lead.x+2 && k.y == lead.y {
				k.x -= 1
			} else {
				k.x = prevLead.x
				k.y = prevLead.y
			}
		} else if k.x == prevLead.x+1 && k.y == prevLead.y-1 {
			/*
				.........
				..MHH....
				..HOH....
				..THM....
				.........
			*/
			if k.x == lead.x && k.y == lead.y-2 {
				k.y += 1
			} else if k.x == lead.x+2 && k.y == lead.y {
				k.x -= 1
			} else {
				k.x = prevLead.x
				k.y = prevLead.y
			}
		} else if k.x == prevLead.x-1 && k.y == prevLead.y-1 {
			/*
				.........
				..THM....
				..HOH....
				..MHH....
				.........
			*/
			if k.x == lead.x && k.y == lead.y-2 {
				k.y += 1
			} else if k.x == lead.x-2 && k.y == lead.y {
				k.x += 1
			} else {
				k.x = prevLead.x
				k.y = prevLead.y
			}
		} else {
			k.x = prevLead.x
			k.y = prevLead.y
		}
	}
}

func (r rope) drawTail(visited [][]bool) {
	visited[r.tail().x][r.tail().y] = true
}

func (r rope) tail() knot {
	return r.knots[r.len()-1]
}

func (r rope) len() int {
	return len(r.knots)
}

func readDirections(filename string) []direction {
	dirs := []direction{}
	file.ParseFile(filename, func(s string) int {
		arr := strings.Split(s, " ")
		dir := parseDirection(arr[0])
		n, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}
		for i := 0; i < n; i++ {
			dirs = append(dirs, dir)
		}
		return -1
	})
	return dirs
}

func parseDirection(s string) direction {
	switch s {
	case "U":
		return up
	case "D":
		return down
	case "L":
		return left
	case "R":
		return right
	}
	panic(fmt.Sprintf("parseDirection: can't parse direction from %s", s))
}

const SCREEN_EMPTY_RUNE = '.'

func (r rope) draw(screen [][]rune) {
	for i, knot := range r.knots {
		screen[knot.x][knot.y] = rune(strconv.Itoa(i)[0])
	}
}

func (r rope) clear(screen [][]rune) {
	for _, knot := range r.knots {
		screen[knot.x][knot.y] = SCREEN_EMPTY_RUNE
	}
}
