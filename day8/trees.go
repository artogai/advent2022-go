package day8

import (
	"advent2022/file"
	"strconv"
)

func CountVisible(filename string) int {
	heights := readHeights(filename)
	visible := make2d[bool](len(heights))

	for i := 0; i < 4; i++ {
		checkFromLeft(heights, visible, i)
		rotate(heights)
	}

	cnt := 0
	for _, row := range visible {
		for _, isVisible := range row {
			if isVisible {
				cnt += 1
			}
		}
	}

	return cnt
}

func FindMaxScore(filename string) int {
	heights := readHeights(filename)
	scores := make2d[int](len(heights))

	for i := 0; i < 4; i++ {
		calcScoresFromLeft(heights, scores, i)
		rotate(heights)
	}

	maxScore := 0
	for _, row := range scores {
		for _, score := range row {
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func calcScoresFromLeft(heights [][]int, scores [][]int, rotationN int) {
	n := len(heights)
	for i, row := range heights {
		for j, height := range row {
			rI, rJ := restoreCoords(i, j, n, rotationN)
			if !(rI == 0 || rJ == 0 || rI == n-1 || rJ == n-1) {
				score := calcScore(j, height, row)
				if scores[rI][rJ] == 0 {
					scores[rI][rJ] = score
				} else {
					scores[rI][rJ] = scores[rI][rJ] * score
				}
			}
		}
	}
}

func calcScore(i int, height int, row []int) int {
	score := 0
	for c := i - 1; c >= 0; c-- {
		if row[c] >= height {
			return score + 1
		} else {
			score += 1
		}
	}
	return score
}

func checkFromLeft(heights [][]int, visible [][]bool, rotationN int) {
	n := len(heights)
	for i, row := range heights {
		maxHeight := -1
		for j, height := range row {
			if height > maxHeight {
				rI, rJ := restoreCoords(i, j, n, rotationN)
				visible[rI][rJ] = true
				maxHeight = height
			}
		}
	}
}

func restoreCoords(i, j, n int, rotationN int) (int, int) {
	rI := i
	rJ := j

	for k := 0; k < rotationN; k++ {
		temp := rI
		rI = rJ
		rJ = n - temp - 1
	}

	return rI, rJ
}

func readHeights(filename string) [][]int {
	lines := file.ReadFileLines(filename)
	heights := make([][]int, len(lines))

	for i, line := range file.ReadFileLines(filename) {
		heights[i] = make([]int, len(line))
		for j, str := range line {
			height, _ := strconv.Atoi(string(str))
			heights[i][j] = height
		}
	}
	return heights
}

// rotate matrix anti-clockwise O(n2)
func rotate[A any](m [][]A) {
	n := len(m)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := m[i][j]
			m[i][j] = m[j][n-1-i]
			m[j][n-1-i] = m[n-1-i][n-1-j]
			m[n-1-i][n-1-j] = m[n-1-j][i]
			m[n-1-j][i] = temp
		}
	}
}

func make2d[A any](n int) [][]A {
	newArr := make([][]A, n)
	for i := range newArr {
		newArr[i] = make([]A, n)
	}
	return newArr
}
