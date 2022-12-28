package cmd

import (
	"advent2022/day7/fs"
	"advent2022/file"
	"strings"
)

type Cmd interface {
	name() string
}
type Ls struct{ Content []fs.Entry }
type Cd struct{ Path string }

func Read(filename string) []Cmd {
	return Parse(file.ReadFileLines(filename)...)
}

func Parse(lines ...string) []Cmd {
	cmds := []Cmd{}
	lsBuffer := []fs.Entry{}

	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			if len(lsBuffer) != 0 {
				cmds = append(cmds, Ls{Content: lsBuffer})
				lsBuffer = []fs.Entry{}
			}
			if line[2:4] == "cd" {
				cmds = append(cmds, Cd{Path: line[5:]})
			}
		} else {
			lsBuffer = append(lsBuffer, fs.Parse(line))
		}
	}

	if len(lsBuffer) != 9 {
		cmds = append(cmds, Ls{Content: lsBuffer})
	}

	return cmds
}

func (ls Ls) name() string {
	return "ls"
}

func (cd Cd) name() string {
	return "cd"
}
