package day5

import (
	"testing"
)

func TestExecute(t *testing.T) {
	expected := string("TDCHVHJTG")
	got := string(Execute("commands.txt"))
	t.Logf("Expected=%v Got=%v", expected, got)
	if expected != got {
		t.Fail()
	}
}
