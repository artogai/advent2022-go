package day18

type cube struct{ x, y, z int }

func surfaceArea(cubes []cube) int {
	m := map[cube]struct{}{}
	for _, c := range cubes {
		m[c] = struct{}{}
	}

	area := 0
	for _, c := range cubes {
		if _, exists := m[cube{c.x + 1, c.y, c.z}]; !exists {
			area++
		}
		if _, exists := m[cube{c.x - 1, c.y, c.z}]; !exists {
			area++
		}
		if _, exists := m[cube{c.x, c.y + 1, c.z}]; !exists {
			area++
		}
		if _, exists := m[cube{c.x, c.y - 1, c.z}]; !exists {
			area++
		}
		if _, exists := m[cube{c.x, c.y, c.z + 1}]; !exists {
			area++
		}
		if _, exists := m[cube{c.x, c.y, c.z - 1}]; !exists {
			area++
		}
	}

	return area
}

func surfaceAreaFixed(cubes []cube, n int) int {
	s := make([][][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			s[i][j] = make([]int, n)
		}
	}

	for _, c := range cubes {
		s[c.x][c.y][c.z] = 2
	}

	fillWithWater(s, n)

	area := 0
	for _, c := range cubes {
		if c.x+1 < n && s[c.x+1][c.y][c.z] == 1 {
			area++
		}
		if s[c.x-1][c.y][c.z] == 1 {
			area++
		}
		if c.y+1 < n && s[c.x][c.y+1][c.z] == 1 {
			area++
		}
		if s[c.x][c.y-1][c.z] == 1 {
			area++
		}
		if c.z+1 < n && s[c.x][c.y][c.z+1] == 1 {
			area++
		}
		if s[c.x][c.y][c.z-1] == 1 {
			area++
		}
	}

	return area
}

func fillWithWater(s [][][]int, n int) {
	var rec func(int, int, int)
	rec = func(x, y, z int) {
		if x < 0 || x >= n || y < 0 || y >= n || z < 0 || z >= n || s[x][y][z] != 0 {
			return
		}
		s[x][y][z] = 1
		rec(x+1, y, z)
		rec(x-1, y, z)
		rec(x, y+1, z)
		rec(x, y-1, z)
		rec(x, y, z+1)
		rec(x, y, z-1)
	}
	rec(0, 0, 0)
}
