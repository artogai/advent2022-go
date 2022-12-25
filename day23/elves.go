package day23

import (
	"fmt"
	"math"
)

type direction int

const (
	N direction = iota
	S direction = iota
	W direction = iota
	E direction = iota
)

type coordinate struct{ row, column int }

func Move(elves map[coordinate]struct{}, n int) map[coordinate]struct{} {
	for i := 0; i < n; i++ {
		dir := direction(i % 4)
		elves, _ = move(elves, dir)
	}
	return elves
}

func MoveUntilFinished(elves map[coordinate]struct{}) (map[coordinate]struct{}, int) {
	finished := false
	for i := 0; ; i++ {
		dir := direction(i % 4)
		elves, finished = move(elves, dir)
		if finished {
			return elves, i + 1
		}
	}
}

func CountFreeSpace(elves map[coordinate]struct{}) int {
	minX, maxX, minY, maxY := findBounds(elves)
	return countFreeSpace(elves, minX, maxX, minY, maxY)
}

func countFreeSpace(elves map[coordinate]struct{}, minX, maxX, minY, maxY int) int {
	return (maxX-minX+1)*(maxY-minY+1) - len(elves)
}

func findBounds(elves map[coordinate]struct{}) (int, int, int, int) {
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for e := range elves {
		if e.row < minX {
			minX = e.row
		}
		if e.row > maxX {
			maxX = e.row
		}
		if e.column < minY {
			minY = e.column
		}
		if e.column > maxY {
			maxY = e.column
		}
	}
	return minX, maxX, minY, maxY
}

func move(elves map[coordinate]struct{}, dir direction) (map[coordinate]struct{}, bool) {
	moves, finished := considerMoves(elves, dir)
	moves = deduplicate(moves)
	elves = applyMoves(elves, moves)
	return elves, finished
}

func applyMoves(elves map[coordinate]struct{}, moves map[coordinate]coordinate) map[coordinate]struct{} {
	nextElves := map[coordinate]struct{}{}
	for e := range elves {
		nextElves[e] = struct{}{}
	}
	for e, coord := range moves {
		delete(nextElves, e)
		nextElves[coord] = struct{}{}
	}
	return nextElves
}

func deduplicate(moves map[coordinate]coordinate) map[coordinate]coordinate {
	seen := map[coordinate]int{}
	unique := map[coordinate]coordinate{}
	for _, coord := range moves {
		seen[coord] += 1
	}
	for e, coord := range moves {
		if seenCnt := seen[coord]; seenCnt == 1 {
			unique[e] = coord
		}
	}
	return unique
}

func considerMoves(elves map[coordinate]struct{}, dir direction) (map[coordinate]coordinate, bool) {
	finished := true
	moves := map[coordinate]coordinate{}
	for e := range elves {
		var pos coordinate
		posSet := false
		noOneAround := true

		for i := 0; i < 4; i++ {
			dir := direction((int(dir) + i) % 4)
			p, free := considerMove(e, dir, elves)
			if !free {
				noOneAround = false
				finished = false
			} else if !posSet {
				pos = p
				posSet = true
			}
		}

		if !noOneAround && posSet {
			moves[e] = pos
		}
	}
	return moves, finished
}

func considerMove(c coordinate, dir direction, elves map[coordinate]struct{}) (coordinate, bool) {
	switch dir {
	case N:
		return considerMoveAt([]coordinate{{c.row - 1, c.column}, {c.row - 1, c.column - 1}, {c.row - 1, c.column + 1}}, elves)
	case S:
		return considerMoveAt([]coordinate{{c.row + 1, c.column}, {c.row + 1, c.column - 1}, {c.row + 1, c.column + 1}}, elves)
	case W:
		return considerMoveAt([]coordinate{{c.row, c.column - 1}, {c.row - 1, c.column - 1}, {c.row + 1, c.column - 1}}, elves)
	case E:
		return considerMoveAt([]coordinate{{c.row, c.column + 1}, {c.row - 1, c.column + 1}, {c.row + 1, c.column + 1}}, elves)
	}
	panic(fmt.Sprint("Unknown direction: ", dir))
}

func considerMoveAt(points []coordinate, elves map[coordinate]struct{}) (coordinate, bool) {
	for _, p := range points {
		if _, exists := elves[p]; exists {
			return coordinate{}, false
		}
	}
	return points[0], true
}

func print(elves map[coordinate]struct{}) {
	minX, maxX, minY, maxY := findBounds(elves)
	for r := minX; r <= maxX; r++ {
		for c := minY; c <= maxY; c++ {
			if _, exists := elves[coordinate{r, c}]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
