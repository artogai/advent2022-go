package day5

import (
	"testing"
)

func TestExecute(t *testing.T) {
	s := generateStacks()
	s.ExecuteCommandsFromFile("commands.txt", false)

	expected := string("TDCHVHJTG")
	got := string(s.ListTop())

	t.Logf("Expected=%v Got=%v", expected, got)
	if expected != got {
		t.Fail()
	}
}

func TestExecuteNewCrane(t *testing.T) {
	s := generateStacks()
	s.ExecuteCommandsFromFile("commands.txt", true)

	expectedNew := string("NGCMPJLHV")
	gotNew := string(s.ListTop())

	t.Logf("Expected=%v Got=%v", expectedNew, gotNew)
	if expectedNew != gotNew {
		t.Fail()
	}
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
