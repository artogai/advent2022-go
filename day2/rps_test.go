package day2

import "testing"

func TestPlayTournament(t *testing.T) {
	expected := 15691
	got := PlayTournament("strategy.txt")
	t.Logf("Expected=%d Got=%d", expected, got)
	if got != expected {
		t.Fail()
	}
}

func TestPlayPlannedTournament(t *testing.T) {
	expected := 12989
	got := PlayPlannedTournament("strategy.txt")
	t.Logf("Expected=%d Got=%d", expected, got)
	if got != expected {
		t.Fail()
	}
}
