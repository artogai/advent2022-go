package day17

import (
	"advent2022/file"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCave2022(t *testing.T) {
	jets := readJets("jets.txt")
	c := newCave(7, 5000, []jetDir(jets))
	c.drop(0, 2022)
	require.Equal(t, 3193, c.towerHeight())
}

func TestCaveTrillion(t *testing.T) {
	n := 1000000000000
	jets := readJets("jets.txt")
	c := newCave(7, 10000, []jetDir(jets))
	c.dropSkipping(n)
	require.Equal(t, 1577650429835, c.towerHeight())
}

func readJets(filename string) []jetDir {
	jetsChars := []rune(file.ReadFile(filename))
	jets := make([]jetDir, len(jetsChars))
	for i, r := range jetsChars {
		switch r {
		case '>':
			jets[i] = true
		case '<':
			jets[i] = false
		default:
			panic("invalid jet direction")
		}
	}
	return jets
}
