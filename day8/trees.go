package day8

import (
	"advent2022/file"
	"advent2022/matrix"
	"strconv"
)

func CountVisible(filename string) int {
	heights := readHeights(filename)
	visible := matrix.NewSquare[bool](heights.Size())

	for i := 0; i < 4; i++ {
		checkFromLeft(heights, visible, i)
		heights.RotateAntiClockwise()
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

func MaxScore(filename string) int {
	heights := readHeights(filename)
	scores := matrix.NewSquare[int](heights.Size())

	for i := 0; i < 4; i++ {
		calcScoresFromLeft(heights, scores, i)
		heights.RotateAntiClockwise()
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

func calcScoresFromLeft(heights matrix.Square[int], scores matrix.Square[int], rotationN int) {
	n := heights.Size()
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

func readHeights(filename string) matrix.Square[int] {
	lines := file.ReadFileLines(filename)
	heights := matrix.NewSquare[int](len(lines))
	for i, line := range lines {
		for j, str := range line {
			height, err := strconv.Atoi(string(str))
			if err != nil {
				panic(err)
			}
			heights[i][j] = height
		}
	}
	return heights
}
