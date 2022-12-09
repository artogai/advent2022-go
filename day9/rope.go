package day9

import (
	"advent2022/file"
	"fmt"
	"strconv"
	"strings"
)

func CountRopeVisits(filename string, fieldSize int) int {
	dirs := readDirections(filename)
	field := make([][]bool, fieldSize)
	for i := range field {
		field[i] = make([]bool, fieldSize)
	}

	simulateRope(field, dirs)

	cnt := 0
	for i := range field {
		for j := range field[i] {
			if field[i][j] {
				cnt += 1
			}
		}
	}
	return cnt
}

type direction int

const (
	Up    direction = iota
	Down            = iota
	Left            = iota
	Right           = iota
)

func simulateRope(field [][]bool, dirs []direction) {
	headX := len(field) / 2
	headY := headX
	tailX := headX
	tailY := headX

	for _, dir := range dirs {
		switch dir {
		case Up:
			if tailX > headX {
				tailX = headX
				tailY = headY
			}
			headX -= 1
		case Down:
			if tailX < headX {
				tailX = headX
				tailY = headY
			}
			headX += 1
		case Left:
			if tailY > headY {
				tailX = headX
				tailY = headY
			}
			headY -= 1
		case Right:
			if tailY < headY {
				tailX = headX
				tailY = headY
			}
			headY += 1
		}
		field[tailX][tailY] = true
	}
}

func readDirections(filename string) []direction {
	return parseDirections(file.ReadLines(filename)...)
}

func parseDirections(strs ...string) []direction {
	dirs := []direction{}
	for _, s := range strs {
		arr := strings.Split(s, " ")
		dir := parseDirection(arr[0])
		n, _ := strconv.Atoi(arr[1])
		for i := 0; i < n; i++ {
			dirs = append(dirs, dir)
		}
	}
	return dirs
}

func parseDirection(s string) direction {
	switch s {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right
	}
	panic(fmt.Sprintf("can't parse direction from %s", s))
}
