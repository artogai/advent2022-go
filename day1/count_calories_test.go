package day1

import "testing"

func TestCountCaloriesSample(t *testing.T) {
	countCaloriesTest(15, "inventories_sample.txt", t)
}

func TestCountCalories(t *testing.T) {
	countCaloriesTest(69693, "inventories.txt", t)
}

func countCaloriesTest(expected int, filename string, t *testing.T) {
	got := CountCalories(filename)
	t.Logf("Expected=%d Got=%d", expected, got)
	if got != expected {
		t.Fail()
	}
}
