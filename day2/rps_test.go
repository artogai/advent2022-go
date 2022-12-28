package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "strategy.txt"

func TestPlayTournament(t *testing.T) {
	require.Equal(t, 15691, PlayTournament(path))
}

func TestPlayPlannedTournament(t *testing.T) {
	require.Equal(t, 12989, PlayPlannedTournament(path))
}
