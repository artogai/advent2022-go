package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "rope.txt"

func TestSimRope(t *testing.T) {
	require.Equal(t, 6642, CalcVisited(path, 2, 1000))
}

func TestSimRope10(t *testing.T) {
	require.Equal(t, 2765, CalcVisited(path, 10, 1000))
}
