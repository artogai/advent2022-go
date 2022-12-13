package day12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHills(t *testing.T) {
	start, finish, hs := readHeights("hills.txt")
	require.Equal(t, 440, findShortestPath(start, finish, hs))
}
