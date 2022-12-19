package day17

type point struct{ x, y int }

type shape interface {
	points() []point
}

type hline []point
type cross []point
type lshape []point
type vline []point
type square []point

// .C###.
func newHLine(top int) *hline {
	x := 2
	y := top - 4
	points := make([]point, 4)
	for i := 0; i < 4; i++ {
		points[i] = point{x + i, y}
	}
	res := hline(points)
	return &res
}

// ..#..
// .#C#.
// ..#..
func newCross(top int) *cross {
	centerX := 3
	centerY := top - 5
	points := []point{
		{centerX, centerY},
		{centerX - 1, centerY},
		{centerX + 1, centerY},
		{centerX, centerY - 1},
		{centerX, centerY + 1},
	}
	res := cross(points)
	return &res
}

// ...#.
// ...#.
// .##C.
func newLShape(top int) *lshape {
	x := 4
	y := top - 4
	points := []point{
		{x, y},
		{x - 1, y},
		{x - 2, y},
		{x, y - 1},
		{x, y - 2},
	}
	res := lshape(points)
	return &res
}

// .#.
// .#.
// .#.
// .C.
func newVLine(top int) *vline {
	x := 2
	y := top - 4
	points := make([]point, 4)
	for i := 0; i < 4; i++ {
		points[i] = point{x, y - i}
	}
	res := vline(points)
	return &res
}

// .##.
// .C#.
func newSquare(top int) *square {
	x := 2
	y := top - 4
	points := []point{
		{x, y},
		{x + 1, y},
		{x, y - 1},
		{x + 1, y - 1},
	}
	res := square(points)
	return &res
}

func (s *hline) points() []point {
	return *s
}

func (s *cross) points() []point {
	return *s
}

func (s *lshape) points() []point {
	return *s
}

func (s *vline) points() []point {
	return *s
}

func (s *square) points() []point {
	return *s
}

func move(points []point, dx, dy int, field [][]rune) ([]point, bool) {
	newPoints := make([]point, len(points))
	for i, p := range points {
		newPoints[i] = point{p.x + dx, p.y + dy}
	}
	return newPoints, isValid(newPoints, field)
}

func isValid(points []point, field [][]rune) bool {
	for _, p := range points {
		if p.x < 0 || p.x >= len(field[0]) || p.y < 0 || p.y >= len(field) {
			return false
		}
		if field[p.y][p.x] != '.' {
			return false
		}
	}
	return true
}
