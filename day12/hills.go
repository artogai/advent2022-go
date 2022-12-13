package day12

import (
	"advent2022/file"
	"math"
)

type point struct{ x, y int }
type heights [][]int

func findShortestPath(start, finish point, heights heights) int {
	n := len(heights)
	m := len(heights[0])

	distances := make([][]int, n)
	for i := 0; i < n; i++ {
		distances[i] = make([]int, m)
		for j := 0; j < m; j++ {
			distances[i][j] = math.MaxInt / 2
		}
	}

	var move func(point, int)

	tryMove := func(from, to point, dist int) {
		if to.x >= 0 && to.y >= 0 && to.x < n && to.y < m {
			if heights[from.x][from.y]+1 >= heights[to.x][to.y] {
				if dist < distances[to.x][to.y] {
					move(to, dist)
				}
			}
		}
	}

	move = func(cur point, dist int) {
		distances[cur.x][cur.y] = dist
		if cur != finish {
			tryMove(cur, point{cur.x + 1, cur.y}, dist+1)
			tryMove(cur, point{cur.x - 1, cur.y}, dist+1)
			tryMove(cur, point{cur.x, cur.y + 1}, dist+1)
			tryMove(cur, point{cur.x, cur.y - 1}, dist+1)
		}
	}

	move(start, 0)
	return distances[finish.x][finish.y]
}

func readHeights(filename string) (point, point, heights) {
	lines := file.ReadLines(filename)
	heights := make([][]int, len(lines))
	var s, e point
	for i, line := range file.ReadLines(filename) {
		heights[i] = make([]int, len(line))
		for j, ch := range line {
			switch ch {
			case 'S':
				s = point{i, j}
				heights[i][j] = 0
			case 'E':
				e = point{i, j}
				heights[i][j] = int('z' - 'a')
			default:
				heights[i][j] = int(ch - 'a')
			}
		}
	}
	return s, e, heights
}
