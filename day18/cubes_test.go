package day18

import (
	"advent2022/file"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSurfaceArea(t *testing.T) {
	cubes := parseCubes("cubes.txt")
	area := surfaceArea(cubes)
	require.Equal(t, 3576, area)
}

func TestSurfaceAreaFixed(t *testing.T) {
	cubes := parseCubes("cubes.txt")
	area := surfaceAreaFixed(cubes, 22)
	require.Equal(t, 2066, area)
}

func parseCubes(filename string) []cube {
	lines := file.ReadFileLines(filename)
	cubes := make([]cube, 0, len(lines))
	for _, cStr := range lines {
		arr := strings.Split(cStr, ",")
		x, err := strconv.Atoi(arr[0])
		if err != nil {
			panic("invalid x")
		}
		y, err := strconv.Atoi(arr[1])
		if err != nil {
			panic("invalid y")
		}
		z, err := strconv.Atoi(arr[2])
		if err != nil {
			panic("invalid z")
		}
		cubes = append(cubes, cube{x + 1, y + 1, z + 1})
	}
	return cubes
}
