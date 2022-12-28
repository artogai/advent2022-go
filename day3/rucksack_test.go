package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "rucksacks.txt"

func TestMissplacedItemsScore(t *testing.T) {
	require.Equal(t, 8072, MissplacedItemsScore(path))
}

func TestBadgesScore(t *testing.T) {
	require.Equal(t, 2567, BadgesScore(path))
}
