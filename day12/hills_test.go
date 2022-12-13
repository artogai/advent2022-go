package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindShortestPath(t *testing.T) {
	start, finish, hs := readHeights("hills.txt")
	require.Equal(t, 440, findShortestPath(start, finish, hs))
}

func TestFindShortestPathAnyStart(t *testing.T) {
	_, finish, hs := readHeights("hills.txt")
	paths := []int{}
	for i := 0; i < len(hs); i++ {
		for j := 0; j < len(hs[0]); j++ {
			if hs[i][j] == 0 {
				paths = append(paths, findShortestPath(point{i, j}, finish, hs))
			}
		}
	}

	min := paths[0]
	for _, p := range paths {
		if p < min {
			min = p
		}
	}

	require.Equal(t, 439, min)
}
