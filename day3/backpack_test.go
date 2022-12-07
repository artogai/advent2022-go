package day3

import "testing"

func TestCountMissplacedItemsScore(t *testing.T) {
	expected := 8072
	got := CountMissplacedItemsScore("backpacks.txt")
	t.Logf("Expected=%d Got=%d", expected, got)
	if got != expected {
		t.Fail()
	}
}
