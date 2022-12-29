package instruction

import (
	"advent2022/file"
	"strconv"
	"strings"
)

type Instruction interface{ isInstruction() bool }
type Noop struct{}
type Addx struct{ Value int }

func Read(filename string) []Instruction {
	return file.ParseFile(filename, parse)
}

func parse(s string) Instruction {
	if s == "noop" {
		return Noop{}
	} else {
		arr := strings.Split(s, " ")
		value, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}
		return Addx{value}
	}
}

func (Noop) isInstruction() bool { return true }
func (Addx) isInstruction() bool { return true }
