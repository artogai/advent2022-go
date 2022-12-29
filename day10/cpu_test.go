package day10

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var path = "instructions.txt"

func TestTakeCpuMeasurements(t *testing.T) {
	cycles := []int{20, 60, 100, 140, 180, 220}
	assert.Equal(t, 10760, TakeCpuMeasurements(path, cycles))
}

func TestDrawCrt(t *testing.T) {
	t.SkipNow()
	res := DrawCRT(path, 40)
	fmt.Println(res)
}
