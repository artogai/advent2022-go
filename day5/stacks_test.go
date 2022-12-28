package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "commands.txt"

func TestExecute(t *testing.T) {
	s := generateStacks()
	cmds := ReadCommands(path)
	s.Execute(cmds, false)
	require.EqualValues(t, "TDCHVHJTG", s.ListTop())
}

func TestExecuteNewCrane(t *testing.T) {
	s := generateStacks()
	cmds := ReadCommands(path)
	s.Execute(cmds, true)
	require.EqualValues(t, "NGCMPJLHV", s.ListTop())
}

func generateStacks() Stacks {
	return Stacks{
		Stack[rune]{'J', 'H', 'G', 'M', 'Z', 'N', 'T', 'F'},
		Stack[rune]{'V', 'W', 'J'},
		Stack[rune]{'G', 'V', 'L', 'J', 'B', 'T', 'H'},
		Stack[rune]{'B', 'P', 'J', 'N', 'C', 'D', 'V', 'L'},
		Stack[rune]{'F', 'W', 'S', 'M', 'P', 'R', 'G'},
		Stack[rune]{'G', 'H', 'C', 'F', 'B', 'N', 'V', 'M'},
		Stack[rune]{'D', 'H', 'G', 'M', 'R'},
		Stack[rune]{'H', 'N', 'M', 'V', 'Z', 'D'},
		Stack[rune]{'G', 'N', 'F', 'H'},
	}
}
