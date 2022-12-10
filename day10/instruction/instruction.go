package instruction

import (
	"advent2022/file"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type instruction interface {
	name() string
}
type Noop struct{}
type Addx struct{ Value int }

func Read(filename string) []instruction {
	return lo.Map(file.ReadLines(filename), func(line string, _ int) instruction {
		return parse(line)
	})
}

func parse(s string) instruction {
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

func (Noop) name() string {
	return "noop"
}

func (Addx) name() string {
	return "addx"
}
