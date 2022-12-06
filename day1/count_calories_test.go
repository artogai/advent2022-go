package day1

import "testing"

func TestCountCaloriesSample(t *testing.T) {
	countCaloriesTest(15, "inventories_sample.txt", 1, t)
}

func TestCountCaloriesSampleTop2(t *testing.T) {
	countCaloriesTest(22, "inventories_sample.txt", 2, t)
}

func TestCountCalories(t *testing.T) {
	countCaloriesTest(69693, "inventories.txt", 1, t)
}

func TestCountCaloriesTop3(t *testing.T) {
	countCaloriesTest(200945, "inventories.txt", 3, t)
}

func countCaloriesTest(expected int, filename string, n int, t *testing.T) {
	got, err := CountCaloriesTopN(filename, n)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Expected=%d Got=%d", expected, got)
	if got != expected {
		t.Fail()
	}
}
