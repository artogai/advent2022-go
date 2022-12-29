package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	sampleTrees = "trees_sample.txt"
	trees       = "trees.txt"
)

func TestCountVisible(t *testing.T) {
	require.Equal(t, 21, CountVisible(sampleTrees))
	require.Equal(t, 1818, CountVisible(trees))
}

func TestFindMaxScore(t *testing.T) {
	require.Equal(t, 8, MaxScore(sampleTrees))
	require.Equal(t, 368368, MaxScore(trees))
}
