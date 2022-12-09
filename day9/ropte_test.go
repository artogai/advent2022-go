package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRopeSim(t *testing.T) {
	require.Equal(t, 6642, CountRopeVisits("rope.txt", 1000))
}
