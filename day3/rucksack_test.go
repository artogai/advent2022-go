package day3

import "testing"

func TestCalcMissplacedItemsScore(t *testing.T) {
	expected := 8072
	got := CalcMissplacedItemsScore("rucksacks.txt")
	t.Logf("Expected=%d Got=%d", expected, got)
	if got != expected {
		t.Fail()
	}
}

func TestBadgesScore(t *testing.T) {
	expected := 8072
	got := CalcBadgesScore("rucksacks.txt")
	t.Logf("Expected=%d Got=%d", expected, got)
	if got != expected {
		t.Fail()
	}
}
