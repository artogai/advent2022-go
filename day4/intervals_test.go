package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountIntervalsFullyContains(t *testing.T) {
	require.Equal(t, 651, CountIntervalsFullyContains("intervals.txt"))
}

func TestCountIntervalsOverlap(t *testing.T) {
	require.Equal(t, 956, CountIntervalsOverlap("intervals.txt"))
}
