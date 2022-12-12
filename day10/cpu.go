package day10

import (
	"advent2022/day10/instruction"
	"strings"

	"github.com/samber/lo"
)

func TakeCpuMeasurements(filename string, cycles []int) int {
	measurements := make([]int, 0, len(cycles))
	instructions := instruction.Read(filename)
	cpu(instructions, func(cycle, register int) {
		if len(cycles) > 0 {
			if cycle == cycles[0] {
				cycles = cycles[1:]
				measurements = append(measurements, cycle*register)
			}
		}
	})
	return lo.Sum(measurements)
}

func DrawCRT(filename string) string {
	resolution := 40
	instructions := instruction.Read(filename)
	line := make([]rune, resolution)
	clear(line, resolution)
	var str strings.Builder

	cpu(instructions, func(cycle, register int) {
		if cycle != 1 && cycle%resolution == 1 {
			str.WriteString(string(line))
			str.WriteString("\n")
			clear(line, resolution)
		}
		draw(line, cycle, register, resolution)
	})
	return str.String()
}

func cpu(insts []instruction.Instruction, onCycle func(cycle, register int)) {
	cycle := 1
	register := 1
	for _, inst := range insts {
		switch instT := inst.(type) {
		case instruction.Noop:
			onCycle(cycle, register)
			cycle += 1
		case instruction.Addx:
			onCycle(cycle, register)
			cycle += 1
			onCycle(cycle, register)
			cycle += 1
			register += instT.Value
		}
	}
	onCycle(cycle, register)
}

func clear(line []rune, resolution int) {
	for i := range line {
		line[i] = '.'
	}
}

func draw(line []rune, cycle int, register int, resolution int) {
	if cycle%resolution >= register && cycle%resolution <= register+2 {
		line[cycle%resolution-1] = '#'
	}
}
