package day10

import (
	"advent2022/day10/instruction"
	"strings"

	"github.com/samber/lo"
)

func Cpu2() string {
	instructions := instruction.Read("instructions.txt")
	cycle := 1
	register := 1
	line := make([]rune, 40)
	for i := range line {
		line[i] = '.'
	}
	var str strings.Builder
	for _, inst := range instructions {
		switch instT := inst.(type) {
		case instruction.Noop:
			draw(line, cycle, register)
			cycle += 1
			if cycle%40 == 1 {
				str.WriteString(string(line) + "\n")
				line = make([]rune, 40)
				for i := range line {
					line[i] = '.'
				}
			}
		case instruction.Addx:
			draw(line, cycle, register)
			cycle += 1
			if cycle%40 == 1 {
				str.WriteString(string(line) + "\n")
				line = make([]rune, 40)
				for i := range line {
					line[i] = '.'
				}
			}
			draw(line, cycle, register)
			cycle += 1
			if cycle%40 == 1 {
				str.WriteString(string(line) + "\n")
				line = make([]rune, 40)
				for i := range line {
					line[i] = '.'
				}
			}
			register += instT.Value
		}
	}
	return str.String()
}

func draw(line []rune, cycle int, register int) {
	if cycle%40 >= register && cycle%40 <= register+2 {
		line[cycle%40-1] = '#'
	}
}

func Cpu() int {
	instructions := instruction.Read("instructions.txt")
	cycles := []int{20, 60, 100, 140, 180, 220}
	measurements := make([]int, 0, len(cycles))
	cycle := 1
	register := 1
	for _, inst := range instructions {
		switch instT := inst.(type) {
		case instruction.Noop:
			cycle += 1
			cycles, measurements = measure(cycle, register, cycles, measurements)
		case instruction.Addx:
			cycle += 1
			cycles, measurements = measure(cycle, register, cycles, measurements)
			cycle += 1
			register += instT.Value
			cycles, measurements = measure(cycle, register, cycles, measurements)
		}
	}

	return lo.Sum(measurements)
}

func measure(cycle int, register int, cycles []int, measurements []int) ([]int, []int) {
	if len(cycles) == 0 {
		return cycles, measurements
	}

	if cycle == cycles[0] {
		if len(cycles) > 1 {
			cycles = cycles[1:]
		} else {
			cycles = []int{}
		}
		measurements = append(measurements, cycle*register)
	}
	return cycles, measurements
}
