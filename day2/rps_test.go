package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayTournament(t *testing.T) {
	require.Equal(t, 15691, PlayTournament("strategy.txt"))
}

func TestPlayPlannedTournament(t *testing.T) {
	require.Equal(t, 12989, PlayPlannedTournament("strategy.txt"))
}
