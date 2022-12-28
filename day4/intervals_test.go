package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "intervals.txt"

func TestCountContains(t *testing.T) {
	require.Equal(t, 651, CountContains(path))
}

func TestCountOverlaps(t *testing.T) {
	require.Equal(t, 956, CountOverlaps(path))
}
