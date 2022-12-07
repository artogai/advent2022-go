package day4

import "testing"

func TestCountIntervalsFullyContains(t *testing.T) {
	expected := 651
	got := CountIntervalsFullyContains("intervals.txt")
	t.Logf("Expected=%d Got=%d", expected, got)
	if got != expected {
		t.Fail()
	}
}
