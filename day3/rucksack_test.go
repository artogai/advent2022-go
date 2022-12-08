package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcMissplacedItemsScore(t *testing.T) {
	require.Equal(t, 8072, CalcMissplacedItemsScore("rucksacks.txt"))
}

func TestBadgesScore(t *testing.T) {
	require.Equal(t, 2567, CalcBadgesScore("rucksacks.txt"))
}
