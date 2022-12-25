package day23

import (
	"advent2022/file"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFreeSpaceSample(t *testing.T) {
	elves := parse("elves_sample.txt")
	elves = Move(elves, 10)
	require.Equal(t, 110, CountFreeSpace(elves))
}

func TestFreeSpace(t *testing.T) {
	elves := parse("elves.txt")
	elves = Move(elves, 10)
	require.Equal(t, 4068, CountFreeSpace(elves))
}

func TestMoveUntilStopSample(t *testing.T) {
	elves := parse("elves_sample.txt")
	_, rounds := MoveUntilFinished(elves)
	require.Equal(t, 20, rounds)
}

func TestMoveUntilStop(t *testing.T) {
	elves := parse("elves.txt")
	_, rounds := MoveUntilFinished(elves)
	require.Equal(t, 968, rounds)
}

func parse(filename string) map[coordinate]struct{} {
	elves := map[coordinate]struct{}{}
	for i, line := range file.ReadLines(filename) {
		for j, c := range line {
			if c == '#' {
				elves[coordinate{i, j}] = struct{}{}
			}
		}
	}
	return elves
}
