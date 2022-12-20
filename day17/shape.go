package day17

type point struct{ x, y int }

type shape interface {
	reset(top int)
	points() []point
	findFullLine(c *cave) (int, bool)
}

type hline [4]point
type cross [5]point
type lshape [5]point
type vline [4]point
type square [4]point

func (s *hline) reset(top int) {
	x := 2
	y := top - 4
	for i := 0; i < 4; i++ {
		(*s)[i] = point{x + i, y}
	}
}

func (s *cross) reset(top int) {
	x := 3
	y := top - 5
	(*s)[0] = point{x, y - 1}
	(*s)[1] = point{x, y}
	(*s)[2] = point{x, y + 1}
	(*s)[3] = point{x - 1, y}
	(*s)[4] = point{x + 1, y}
}

func (s *lshape) reset(top int) {
	x := 4
	y := top - 4
	(*s)[0] = point{x, y - 2}
	(*s)[1] = point{x, y - 1}
	(*s)[2] = point{x, y}
	(*s)[3] = point{x - 1, y}
	(*s)[4] = point{x - 2, y}
}

func (s *vline) reset(top int) {
	x := 2
	y := top - 4
	for i := 0; i < 4; i++ {
		(*s)[i] = point{x, y - i}
	}
}

func (s *square) reset(top int) {
	x := 2
	y := top - 4
	(*s)[0] = point{x, y - 1}
	(*s)[1] = point{x + 1, y - 1}
	(*s)[2] = point{x, y}
	(*s)[3] = point{x + 1, y}
}

func (s *hline) findFullLine(c *cave) (int, bool) {
	return findFullLine(s[0].y, c)
}

func (s *cross) findFullLine(c *cave) (int, bool) {
	if y, exists := findFullLine(s[1].y, c); exists {
		return y, exists
	}
	if y, exists := findFullLine(s[2].y, c); exists {
		return y, exists
	}
	return -1, false
}

func (s *lshape) findFullLine(c *cave) (int, bool) {
	if y, exists := findFullLine(s[2].y, c); exists {
		return y, exists
	}
	return -1, false
}

func (s *vline) findFullLine(c *cave) (int, bool) {
	for _, p := range s {
		if y, exists := findFullLine(p.y, c); exists {
			return y, exists
		}
	}
	return -1, false
}

func (s *square) findFullLine(c *cave) (int, bool) {
	if y, exists := findFullLine(s[0].y, c); exists {
		return y, exists
	}
	if y, exists := findFullLine(s[2].y, c); exists {
		return y, exists
	}
	return -1, false
}

func (s *hline) points() []point {
	return (*s)[:]
}

func (s *cross) points() []point {
	return (*s)[:]
}

func (s *lshape) points() []point {
	return (*s)[:]
}

func (s *vline) points() []point {
	return (*s)[:]
}

func (s *square) points() []point {
	return (*s)[:]
}

func move(s shape, dx, dy int, c *cave) bool {
	points := s.points()
	for _, p := range points {
		newX := p.x + dx
		newY := p.y + dy
		if newX < 0 || newX >= c.width || newY < 0 || newY >= c.height {
			return false
		}
		if (*c.field)[newY][newX] {
			return false
		}
	}

	for i := range points {
		points[i].x += dx
		points[i].y += dy
	}
	return true
}

func findFullLine(y int, c *cave) (int, bool) {
	for i := 0; i < c.width; i++ {
		if !(*c.field)[y][i] {
			return -1, false
		}
	}
	return y, true
}

// Number of lines above the top to keep clear.
// Lines above (top - clearAboveTop) can contain garbage blocks from previous iterations,
// but this won't affect the current iteration because blocks always spawn
// at 3 lines above the top
const clearAboveTop = 7
