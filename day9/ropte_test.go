package day9

import (
	"fmt"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestSimRope(t *testing.T) {
	require.Equal(t, 6642, calcRopeTailVisited(t, 2, "rope.txt"))
}

func TestSimRope10(t *testing.T) {
	require.Equal(t, 2765, calcRopeTailVisited(t, 10, "rope.txt"))
}

func calcRopeTailVisited(t *testing.T, ropeSize int, filename string) int {
	fieldSize := 1000
	visited := make([][]bool, fieldSize)

	for i := 0; i < fieldSize; i++ {
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

func drawRope(t *testing.T) {
	fieldSize := 30
	visited := make([][]bool, fieldSize)
	screen := make([][]rune, fieldSize)

	for i := 0; i < fieldSize; i++ {
		visited[i] = make([]bool, fieldSize)
		screen[i] = make([]rune, fieldSize)
		for j := 0; j < fieldSize; j++ {
			screen[i][j] = SCREEN_EMPTY_RUNE
		}
	}

	dirs := readDirections("rope_sample.txt")
	rope := newRope(fieldSize/2, fieldSize/2, 10)

	for _, dir := range dirs {
		rope.traverse(dir, visited, screen)
		fmt.Println(dir.toString())
		fmt.Println(drawScreen(screen))
	}

	cnt := calcVisited(visited)
	fmt.Println(drawVisited(visited))
	fmt.Println(cnt)
}

func drawScreen(screen [][]rune) string {
	var str strings.Builder
	for _, row := range screen {
		str.WriteString(string(row) + "\n")
	}
	return str.String()
}

func drawVisited(visited [][]bool) string {
	var str strings.Builder
	for _, row := range visited {
		s := string(lo.Map(row, func(b bool, _ int) rune {
			if b {
				return '#'
			} else {
				return SCREEN_EMPTY_RUNE
			}
		}))
		str.WriteString(s + "\n")
	}
	return str.String()
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
