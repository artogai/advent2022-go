package day11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMonkeyBusinessSample20Rounds(t *testing.T) {
	ms := read("monkeys_sample.txt")
	ms.runSimulation(20, true)
	require.Equal(t, 10605, ms.businessScore())
}

func TestMonkeyBusiness20Rounds(t *testing.T) {
	ms := read("monkeys.txt")
	ms.runSimulation(20, true)
	require.Equal(t, 110220, ms.businessScore())
}

func TestMonkeyBusinessSample10000Rounds(t *testing.T) {
	ms := read("monkeys.txt")
	ms.runSimulation(10000, false)
	require.Equal(t, 19457438264, ms.businessScore())
}
