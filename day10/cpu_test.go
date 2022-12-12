package day10

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeCpuMeasurements(t *testing.T) {
	cycles := []int{20, 60, 100, 140, 180, 220}
	assert.Equal(t, 10760, TakeCpuMeasurements("instructions.txt", cycles))
}

func TestDrawCrt(t *testing.T) {
	res := DrawCRT("instructions.txt")
	//FPGPHFGH
	fmt.Println(res)
	t.SkipNow()
}
