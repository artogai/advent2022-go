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

// func TestMonkeyBusiness10000Rounds(t *testing.T) {
// 	ms := read("monkeys_sample.txt")
// 	ms.runSimulation(1000, false)
// 	fmt.Println(ms.businessScore())
// 	// require.Equal(t, 37691855600, ms.businessScore())
// }
